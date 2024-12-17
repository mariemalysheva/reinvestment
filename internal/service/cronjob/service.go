package cronjob

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/jackc/pgx/v5"
	contract_reinvestment "github.com/mariemalysheva/tokenized-reinvestment/contracts/reinvestment"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/node"
	repository "github.com/mariemalysheva/tokenized-reinvestment/internal/repo"
	crontab "github.com/robfig/cron/v3"
	"log"
	"time"
)

type Service struct {
	repo     *repository.Repository
	ownerKey *ecdsa.PrivateKey
	node     *node.Client
	entryID  crontab.EntryID
	c        *crontab.Cron
	schedule string
}

func New(
	repo *repository.Repository,
	ownerKey *ecdsa.PrivateKey,
	node *node.Client,
	c *crontab.Cron,
) *Service {
	return &Service{
		repo:     repo,
		ownerKey: ownerKey,
		node:     node,
		c:        c,
	}
}

func (s *Service) ReinvestSavings() error {
	log.Println("reinvestment cronjob started")
	defer log.Println("reinvestment cronjob finished")

	reinvestmentContract, err := s.node.GetReinvestmentContract()
	if err != nil {
		return err
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(s.ownerKey, s.node.GetChainID())
	if err != nil {
		return err
	}

	txHash, err := reinvestmentContract.ReinvestSavings(txOpts)
	if err != nil {
		return err
	}
	endTime := time.Now().UTC()

	log.Println("savings reinvested in tx", txHash.Hash().Hex())

	time.Sleep(15 * time.Second)

	filterer, err := contract_reinvestment.NewContractReinvestmentFilterer(s.node.GetReinvestmentAddress(), s.node.GetNode())
	if err != nil {
		return err
	}

	reinv, err := s.node.GetCurrentReinvestment()
	if err != nil {
		return err
	}

	ctx := context.Background()

	txReceipt, err := s.node.GetNode().TransactionReceipt(ctx, txHash.Hash())
	if err != nil {
		return err
	}

	event, err := filterer.ParseSavingsReinvested(*txReceipt.Logs[2])
	if err != nil {
		return err
	}

	err = s.repo.InTx(ctx, func(ctx context.Context, tx pgx.Tx) (err error) {
		err = s.repo.UpdateReinvestmentPeriodEndTime(ctx, tx, endTime)
		if err != nil {
			return err
		}

		err = s.repo.AddReinvestmentPeriod(ctx, tx, reinv)
		if err != nil {
			return err
		}

		for i, userID := range event.UserIDs {
			bal := event.NewBalances[i].Int64()
			err = s.repo.AddReinvestmentRecord(ctx, tx, userID.Int64(), models.ReinvestmentRecord{
				Reinvestment: models.Reinvestment{ID: reinv.ID},
				Savings:      bal,
				Amount:       bal / reinv.Price,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RescheduleCronHandlerFunc(schedule string, f func() error) (err error) {
	_, err = crontab.ParseStandard(schedule)
	if err != nil {
		return err
	}

	defer log.Printf("reinvestment schedule changed to %s", schedule)

	s.c.Remove(s.entryID)
	return s.AddCronHandlerFunc(schedule, f)
}

func (s *Service) AddCronHandlerFunc(
	schedule string,
	f func() error,
) (err error) {
	if schedule == "" {
		return fmt.Errorf("empty schedule")
	}

	s.schedule = schedule

	s.entryID, err = s.c.AddFunc(schedule, func() {
		inErr := f()
		if inErr != nil {
			log.Printf("err in cronjob handler func: %v", inErr)
		}
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCurrentReinvestmentSchedule() string {
	return s.schedule
}
