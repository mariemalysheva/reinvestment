package admin

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/node"
	repository "github.com/mariemalysheva/tokenized-reinvestment/internal/repo"
	"math/big"
	"time"
)

type Service struct {
	repo *repository.Repository
	priv *ecdsa.PrivateKey
	node *node.Client
}

func New(
	repo *repository.Repository,
	priv *ecdsa.PrivateKey,
	node *node.Client,
) *Service {
	return &Service{
		repo: repo,
		priv: priv,
		node: node,
	}
}

func (s *Service) PostAddUsers(ctx context.Context, users []models.Client) (txHash string, err error) {
	curReinv, err := s.node.GetCurrentReinvestment()
	if err != nil {
		return "", err
	}

	err = s.repo.InTx(ctx, func(ctx context.Context, tx pgx.Tx) (err error) {
		var savings []*big.Int
		var userIDs []*big.Int
		var salt []string

		for i, u := range users {
			users[i], err = s.repo.AddClient(ctx, tx, u)
			if err != nil {
				return err
			}

			users[i].CreatedAt = time.Now()

			users[i].Salt, err = users[i].Sign(s.priv, s.node.GetReinvestmentAddress())
			if err != nil {
				return err
			}

			err = s.repo.AddReinvestmentRecord(ctx, tx, users[i].ID, models.ReinvestmentRecord{
				Reinvestment: models.Reinvestment{ID: curReinv.ID},
				Savings:      users[i].Savings,
				Amount:       users[i].Savings / curReinv.Price,
			})

			fmt.Println(users[i])

			if err != nil {
				return err
			}

			savings = append(savings, big.NewInt(users[i].Savings))
			userIDs = append(userIDs, big.NewInt(users[i].ID))
			salt = append(salt, users[i].Salt)
		}

		reinv, err := s.node.GetReinvestmentContract()
		if err != nil {
			return err
		}

		txOpts, err := s.node.TxOpts(s.priv)
		if err != nil {
			return err
		}

		txAdd, err := reinv.AddUserBatch(txOpts, userIDs, salt, savings)
		if err != nil {
			return err
		}
		txHash = txAdd.Hash().Hex()
		return nil
	})
	return txHash, err
}

func (s *Service) PostTransferUsers(ctx context.Context, users []int64, to common.Address) (txHash string, err error) {

	err = s.repo.InTx(ctx, func(ctx context.Context, tx pgx.Tx) (err error) {
		var userIDs []*big.Int
		for _, userID := range users {
			err = s.repo.DeleteClient(ctx, tx, userID)
			if err != nil {
				return err
			}
			userIDs = append(userIDs, big.NewInt(userID))
		}

		reinv, err := s.node.GetReinvestmentContract()
		if err != nil {
			return err
		}

		txOpts, err := s.node.TxOpts(s.priv)
		if err != nil {
			return err
		}

		txTransfer, err := reinv.TransferUserBatch(txOpts, userIDs, to)
		if err != nil {
			return err
		}
		txHash = txTransfer.Hash().Hex()
		return nil
	})

	return txHash, err
}

func (s *Service) GetUsers(ctx context.Context) (users []models.Client, err error) {
	users, err = s.repo.GetClients(ctx)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (s *Service) GetUserSalt(ctx context.Context, id int64) (user models.Client, err error) {
	user, err = s.repo.GetClient(ctx, id)
	if err != nil {
		return models.Client{}, err
	}

	user.Salt, err = user.Sign(s.priv, s.node.GetReinvestmentAddress())
	if err != nil {
		return models.Client{}, err
	}
	return user, err
}

func (s *Service) GetUserReinvestmentHistory(ctx context.Context, id int64) (hist []models.ReinvestmentRecord, err error) {
	user, err := s.repo.GetClient(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.repo.GetClientReinvestmentHistory(ctx, user.ID)
}

func (s *Service) PostSetRate(ctx context.Context, rate float64) (txHash string, err error) {
	rateInt := int64(rate * models.RateDenom)

	reinv, err := s.node.GetReinvestmentContract()
	if err != nil {
		return "", err
	}

	txOpts, err := s.node.TxOpts(s.priv)
	if err != nil {
		return "", err
	}

	tx, err := reinv.SetRate(txOpts, big.NewInt(rateInt))
	if err != nil {
		return "", err
	}

	err = s.repo.UpdateReinvestmentPeriodRate(ctx, rateInt)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) PostSetPrice(ctx context.Context, price int64) (txHash string, err error) {
	reinv, err := s.node.GetReinvestmentContract()
	if err != nil {
		return "", err
	}

	txOpts, err := s.node.TxOpts(s.priv)
	if err != nil {
		return "", err
	}

	tx, err := reinv.SetAssetPrice(txOpts, big.NewInt(price))
	if err != nil {
		return "", err
	}

	err = s.repo.UpdateReinvestmentPeriodAssetPrice(ctx, price)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) GetReinvestmentPeriods(ctx context.Context) (reinvs []models.Reinvestment, err error) {
	return s.repo.GetReinvestmentPeriods(ctx)
}

func (s *Service) VerifyUserSalt(ctx context.Context, userID int64, salt string) error {
	user, err := s.GetUserSalt(ctx, userID)
	if err != nil {
		return err
	}

	if user.Salt != salt {
		return fmt.Errorf("invalid user salt")
	}

	return nil
}
