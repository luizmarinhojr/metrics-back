package main

import (
	"fmt"
	"os"

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

	
	dirPath := "./docs"
	filePath := "./docs/swagger.json"

	// Valida se o diretório "/docs" existe
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("O diretório '%s' não existe.\n", dirPath)
		} else {
			fmt.Printf("Erro ao verificar o diretório '%s': %v\n", dirPath, err)
		}
		return
	}
	if !dirInfo.IsDir() {
		fmt.Printf("O caminho '%s' não é um diretório.\n", dirPath)
		return
	}
	fmt.Printf("O diretório '%s' existe.\n", dirPath)

	// Valida se o arquivo "/docs/swagger.json" existe
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("O arquivo '%s' não existe.\n", filePath)
		} else {
			fmt.Printf("Erro ao verificar o arquivo '%s': %v\n", filePath, err)
		}
		return
	}
	if fileInfo.IsDir() {
		fmt.Printf("O caminho '%s' é um diretório, não um arquivo.\n", filePath)
		return
	}
	fmt.Printf("O arquivo '%s' existe.\n", filePath)


	api.InitializeApi(dependencies)
}
