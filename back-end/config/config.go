package config

import "os"

var (
	DB_CONN = os.Getenv("DB_CONN")
)
