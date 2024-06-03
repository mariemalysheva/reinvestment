package repository

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
	"time"
)

func (r *Repository) AddReinvestmentPeriod(ctx context.Context, tx pgx.Tx, reinv models.Reinvestment) error {
	const query = `
	insert into reinvestment_period (id, asset_address, reinvestment_rate, asset_price, start_time)
	values ($1, $2, $3, $4, $5)
	on conflict do nothing;
`
	_, err := tx.Exec(ctx, query, reinv.ID, reinv.Asset.Hex(), reinv.Rate, reinv.Price, reinv.Start)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateReinvestmentPeriodEndTime(ctx context.Context, tx pgx.Tx, endTime time.Time) error {
	const query = `
	update reinvestment_period
	set end_time = $1
	where id = (select id from reinvestment_period order by id desc limit 1);
`
	rows, err := tx.Exec(ctx, query, endTime)
	if err != nil {
		if rows.RowsAffected() == 0 {
			return fmt.Errorf("current reinvestment not found")
		}
		return err
	}
	return nil
}

func (r *Repository) UpdateReinvestmentPeriodRate(ctx context.Context, rate int64) error {
	const query = `
	update reinvestment_period
	set reinvestment_rate = $1
	where id = (select id from reinvestment_period order by id desc limit 1);
`
	rows, err := r.db.Exec(ctx, query, rate)
	if err != nil {
		if rows.RowsAffected() == 0 {
			return fmt.Errorf("current reinvestment not found")
		}
		return err
	}
	return nil
}

func (r *Repository) UpdateReinvestmentPeriodAssetPrice(ctx context.Context, price int64) error {
	const query = `
	update reinvestment_period
	set asset_price = $1
	where id = (select id from reinvestment_period order by id desc limit 1);
`
	rows, err := r.db.Exec(ctx, query, price)
	if err != nil {
		if rows.RowsAffected() == 0 {
			return fmt.Errorf("current reinvestment not found")
		}
		return err
	}
	return nil
}

func (r *Repository) AddReinvestmentRecord(ctx context.Context, tx pgx.Tx, userID int64, reinv models.ReinvestmentRecord) error {
	const query = `
	insert into user_reinvestment (client_id, reinvestment_period_id, savings_balance, asset_amount)
	values ($1, $2, $3, $4)
	on conflict do nothing;
`
	_, err := tx.Exec(ctx, query, userID, reinv.Reinvestment.ID, reinv.Savings, reinv.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetReinvestmentPeriods(ctx context.Context) (reinvs []models.Reinvestment, err error) {
	const query = `
	select 
	    rp.id,
	    rp.asset_address,
	    rp.asset_price,
	    rp.reinvestment_rate,
	    rp.start_time,
	    rp.end_time
	
	from public.reinvestment_period rp
	order by id desc;

	
`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		var (
			period    models.Reinvestment
			assetAddr string
		)
		err = rows.Scan(
			&period.ID,
			&assetAddr,
			&period.Price,
			&period.Rate,
			&period.Start,
			&period.End,
		)
		if err != nil {
			return nil, err
		}
		period.Asset = common.HexToAddress(assetAddr)
		reinvs = append(reinvs, period)
	}

	return reinvs, nil
}
