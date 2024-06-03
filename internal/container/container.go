package container

import (
	"context"
	"crypto/ecdsa"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/node"
	repository "github.com/mariemalysheva/tokenized-reinvestment/internal/repo"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/service/admin"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/service/cronjob"
	crontab "github.com/robfig/cron/v3"
)

type Container struct {
	adminSvc *admin.Service
	cronSvc  *cronjob.Service

	// API Implementation
	api *handler.Implementation

	// Repos
	repo    *repository.Repository
	db      *pgxpool.Pool
	node    *node.Client
	key     *ecdsa.PrivateKey
	crontab *crontab.Cron
}

func New() *Container {
	return &Container{}
}

func (c *Container) GetAPI(ctx context.Context) (*handler.Implementation, error) {
	if c.api != nil {
		return c.api, nil
	}

	adminSvc, err := c.GetAdminSvc(ctx)
	if err != nil {
		return nil, err
	}

	cronSvc, err := c.GetCronSvc(ctx)
	if err != nil {
		return nil, err
	}

	c.api = handler.NewImplementation(adminSvc, cronSvc)
	return c.api, nil
}
