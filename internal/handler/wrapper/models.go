package wrapper

import (
	"time"
)

type PostAddUsersReq struct {
	Clients []UserReq `json:"clients"`
}

type PostTransferUsersReq struct {
	ClientIDs []int64 `json:"client_ids"`
	ToAddress string  `json:"to_address"`
}

type PostSetRateReq struct {
	Rate float64 `json:"rate"`
}
type PostSetPriceReq struct {
	Price int64 `json:"price"`
}

type PostSetScheduleReq struct {
	CronSchedule string `json:"cron_schedule"`
}

type UserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Savings   int64  `json:"savings"`
}

type Reinvestment struct {
	ID    int64     `json:"id"`
	Asset string    `json:"asset"`
	Rate  float64   `json:"rate"`
	Price int64     `json:"price"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type ReinvestmentRecord struct {
	Reinvestment Reinvestment `json:"reinvestment"`
	Savings      uint64       `json:"savings"`
	Amount       uint64       `json:"amount"`
}

type TxResp struct {
	TxHash string `json:"tx_hash"`
}
