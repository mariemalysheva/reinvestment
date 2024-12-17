package handler

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
)

type AdminSvc interface {
	PostAddUsers(ctx context.Context, users []models.Client) (txHash string, err error)
	PostTransferUsers(ctx context.Context, users []int64, to common.Address) (txHash string, err error)
	GetUsers(ctx context.Context) (users []models.Client, err error)
	GetUserSalt(ctx context.Context, id int64) (user models.Client, err error)
	GetUserReinvestmentHistory(ctx context.Context, id int64) (hist []models.ReinvestmentRecord, err error)
	PostSetRate(ctx context.Context, price float64) (hash string, err error)
	PostSetPrice(ctx context.Context, price int64) (hash string, err error)
	GetReinvestmentPeriods(ctx context.Context) (reinvs []models.Reinvestment, err error)
	VerifyUserSalt(ctx context.Context, userID int64, salt string) error
}

type SchedulerSvc interface {
	ReinvestSavings() error
	RescheduleCronHandlerFunc(schedule string, f func() error) (err error)
	GetCurrentReinvestmentSchedule() string
}

type Implementation struct {
	adminSvc     AdminSvc
	schedulerSvc SchedulerSvc
}

func NewImplementation(profileSvc AdminSvc, schedulerSvc SchedulerSvc) *Implementation {
	return &Implementation{
		adminSvc:     profileSvc,
		schedulerSvc: schedulerSvc,
	}
}
