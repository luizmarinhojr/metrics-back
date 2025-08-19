package main

import (
	"github.com/luizmarinhojr/metrics/internal/app/dependencies"
	"github.com/luizmarinhojr/metrics/internal/database"
	"github.com/luizmarinhojr/metrics/internal/http/api"
)

// @title Metrics API
// @version 1.0
// @description Esta é uma API de métricas para a aplicação Metrics. A chave da API é armazenada através de cookie, portanto, para conseguir a chave da API basta que faça o login no endpoint "login".
// @externalDocs.url   https://github.com/luizmarinhojr/metrics-back
// @contact.name   Luiz Marinho Support
// @contact.url    https://luizmarinhodev.vercel.app/
// @contact.email  luizmarinhodev@gmail.com
// @host      api.naster.com.br
// @schemes   https
// @BasePath /api/v1
// @SecurityDefinitions.apiKey ApiKeyAuth
// @in cookie
// @name token
func main() {
	db := database.OpenConnection()

	dependencies := dependencies.Inject(db)

	api.InitializeApi(dependencies)
}
