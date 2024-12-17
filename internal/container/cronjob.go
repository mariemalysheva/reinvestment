package container

import (
	"context"
	crontab "github.com/robfig/cron/v3"
)

const (
	everyThirdMonth = "0 0 1 */3 *"
	everyMinute     = "* * * * *"
)

func (c *Container) GetCrontab() *crontab.Cron {
	if c.crontab != nil {
		return c.crontab
	}
	c.crontab = crontab.New()
	return c.crontab

}

func (c *Container) RunCronJobs(
	ctx context.Context,
) error {

	cr := c.GetCrontab()
	defer cr.Start()

	svc, err := c.GetCronSvc(ctx)
	if err != nil {
		return err
	}

	return svc.AddCronHandlerFunc(everyThirdMonth, svc.ReinvestSavings)
}
