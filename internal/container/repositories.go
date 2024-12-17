package container

import (
	"context"
	repository "github.com/mariemalysheva/tokenized-reinvestment/internal/repo"
)

func (c *Container) GetRepo(ctx context.Context) (*repository.Repository, error) {
	if c.repo != nil {
		return c.repo, nil
	}
	db, err := c.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	c.repo = repository.New(db)
	return c.repo, nil
}
