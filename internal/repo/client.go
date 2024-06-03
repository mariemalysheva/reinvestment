package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
)

func (r *Repository) AddClient(ctx context.Context, tx pgx.Tx, user models.Client) (models.Client, error) {
	const query = `
	insert into client (first_name, last_name)
	values ($1, $2)
	returning id
`
	err := tx.QueryRow(ctx, query, user.FirstName, user.LastName).Scan(&user.ID)
	if err != nil {
		return models.Client{}, err
	}

	return user, nil
}

func (r *Repository) DeleteClient(ctx context.Context, tx pgx.Tx, id int64) error {
	const query = `
	delete from client
	where id = $1;
`
	rows, err := tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return fmt.Errorf("user not found by id %d", id)
	}

	const queryReinvestment = `
	delete from user_reinvestment
	where client_id = $1;
`
	_, err = tx.Exec(ctx, queryReinvestment, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetClients(ctx context.Context) (users []models.Client, err error) {
	const query = `
	select id, first_name, last_name, created_at
	from client
`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		user := models.Client{}
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func (r *Repository) GetClient(ctx context.Context, id int64) (user models.Client, err error) {
	const query = `
	select first_name, last_name, created_at
	from client
	where id = $1
`
	err = r.db.QueryRow(ctx, query, id).Scan(&user.FirstName, &user.LastName, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Client{}, fmt.Errorf("client not found by id")
		}
		return models.Client{}, err
	}

	return user, nil
}

func (r *Repository) GetClientReinvestmentHistory(ctx context.Context, id int64) (hist []models.ReinvestmentRecord, err error) {
	const query = `
	select 
	    rp.id,
	    rp.asset_address,
	    rp.asset_price,
	    rp.reinvestment_rate,
	    rp.start_time,
	    rp.end_time,
	    ur.asset_amount,
	    ur.savings_balance
	
	from user_reinvestment ur 
	join public.reinvestment_period rp on rp.id = ur.reinvestment_period_id
	where ur.client_id = $1
	
`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		var (
			rec       models.ReinvestmentRecord
			assetAddr string
		)
		err = rows.Scan(
			&rec.Reinvestment.ID,
			&assetAddr,
			&rec.Reinvestment.Price,
			&rec.Reinvestment.Rate,
			&rec.Reinvestment.Start,
			&rec.Reinvestment.End,
			&rec.Amount,
			&rec.Savings,
		)
		if err != nil {
			return nil, err
		}
		rec.Reinvestment.Asset = common.HexToAddress(assetAddr)
		rec.Reinvestment.Rate = rec.Reinvestment.Rate / models.RateDenom
		hist = append(hist, rec)
	}

	return hist, nil
}
