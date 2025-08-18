package config

import "os"

var (
	DB_CONN        = os.Getenv("DB_CONN")
	DOMAIN         = os.Getenv("DOMAIN_ORIGIN")
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	PATH           = os.Getenv("PATH_COOKIE")
)
