package cronjob

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/jackc/pgx/v5"
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

func (s *Service) ReinvestSavings(ctx context.Context) error {
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

	time.Sleep(20 * time.Second)

	reinv, err := s.node.GetCurrentReinvestment()
	if err != nil {
		return err
	}

	err = s.repo.InTx(ctx, func(ctx context.Context, tx pgx.Tx) (err error) {
		err = s.repo.UpdateReinvestmentPeriodEndTime(ctx, tx, endTime)
		if err != nil {
			return err
		}
		return s.repo.AddReinvestmentPeriod(ctx, tx, reinv)
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RescheduleCronHandlerFunc(
	ctx context.Context, schedule string, f func(ctx context.Context) error) (err error) {
	_, err = crontab.ParseStandard(schedule)
	if err != nil {
		return err
	}

	defer log.Printf("reinvestment schedule changed to %s", schedule)

	s.c.Remove(s.entryID)
	return s.AddCronHandlerFunc(ctx, schedule, f)
}

func (s *Service) AddCronHandlerFunc(
	ctx context.Context,
	schedule string,
	f func(ctx context.Context) error,
) (err error) {
	if schedule == "" {
		return fmt.Errorf("empty schedule")
	}

	s.entryID, err = s.c.AddFunc(schedule, func() {
		inErr := f(ctx)
		if inErr != nil {
			log.Printf("err in cronjob handler func: %v", inErr)
		}
	})
	if err != nil {
		return err
	}
	return nil
}
