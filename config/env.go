package config

import "os"

var (
	DBConn           = os.Getenv("DB_CONN")
	Port             = os.Getenv("PORT")
	RPC              = os.Getenv("RPC")
	OwnerPriv        = os.Getenv("OWNER_PRIV")
	ReinvestmentAddr = os.Getenv("REINVESTMENT_ADDR")
)
