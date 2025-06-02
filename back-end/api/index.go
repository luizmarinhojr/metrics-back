package main

import (
	"github.com/luizmarinhojr/metrics/internal/app/dependencies"
	"github.com/luizmarinhojr/metrics/internal/database"
	"github.com/luizmarinhojr/metrics/internal/http/api"
)

func main() {
	db := database.OpenConnection()

	dependencies := dependencies.Inject(db)

	api.InitializeApi(dependencies)
}
