package container

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/service/admin"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/service/cronjob"
)

func (c *Container) GetAdminSvc(ctx context.Context) (*admin.Service, error) {
	if c.adminSvc != nil {
		return c.adminSvc, nil
	}
	repo, err := c.GetRepo(ctx)
	if err != nil {
		return nil, err
	}

	c.node, err = c.GetNode(ctx)
	if err != nil {
		return nil, err
	}

	c.key, err = c.GetOwnerKey()
	if err != nil {
		return nil, err
	}

	c.adminSvc = admin.New(repo, c.key, c.node)
	return c.adminSvc, nil
}

func (c *Container) GetCronSvc(ctx context.Context) (cronSvc *cronjob.Service, err error) {
	if c.cronSvc != nil {
		return c.cronSvc, nil
	}

	c.crontab = c.GetCrontab()

	c.repo, err = c.GetRepo(ctx)
	if err != nil {
		return nil, err
	}

	c.node, err = c.GetNode(ctx)
	if err != nil {
		return nil, err
	}

	c.key, err = c.GetOwnerKey()
	if err != nil {
		return nil, err
	}

	reinv, err := c.node.GetCurrentReinvestment()
	if err != nil {
		return nil, err
	}

	err = c.repo.InTx(ctx, func(ctx context.Context, tx pgx.Tx) (err error) {
		return c.repo.AddReinvestmentPeriod(ctx, tx, reinv)
	})
	if err != nil {
		return nil, err
	}

	c.cronSvc = cronjob.New(c.repo, c.key, c.node, c.crontab)
	return c.cronSvc, nil
}
