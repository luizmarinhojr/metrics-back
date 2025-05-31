package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/internal/app/dependencies"
)

func InitializeApi(dependencies *dependencies.Dependencies) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},                             // Permite requisições do seu front-end Angular
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // Métodos HTTP permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Headers permitidos
		ExposeHeaders:    []string{"Content-Length"},                                    // Headers que podem ser expostos ao cliente
		AllowCredentials: true,                                                          // Permite o envio de cookies e credenciais
		MaxAge:           300,                                                           // Tempo máximo em segundos que a resposta de pré-voo pode ser armazenada em cache
	}))

	baseURL := r.Group("/api/v1")
	{
		{
			baseURL.GET("brokers", dependencies.Middlewares.CheckAuth.Auth, dependencies.Handlers.BrokerHandler.GetAll)
			baseURL.GET("broker/id", dependencies.Middlewares.CheckAuth.Auth, dependencies.Handlers.BrokerHandler.GetById)
			baseURL.GET("broker/name", dependencies.Middlewares.CheckAuth.Auth, dependencies.Handlers.BrokerHandler.GetByName)
			baseURL.POST("broker", dependencies.Middlewares.CheckAuth.Auth, dependencies.Handlers.BrokerHandler.Create)
		}
		{
			baseURL.GET("metrics", dependencies.Middlewares.CheckAuth.Auth, dependencies.Handlers.MetricHandler.GetAll)
			baseURL.GET("metric/id", dependencies.Middlewares.CheckAuth.Auth, dependencies.Handlers.MetricHandler.GetByID)
			baseURL.POST("metric", dependencies.Middlewares.CheckAuth.Auth, dependencies.Handlers.MetricHandler.Create)
		}
		{
			baseURL.POST("login", dependencies.Handlers.UserHandler.GetByEmail)
			baseURL.POST("signup", dependencies.Handlers.UserHandler.Create)
			baseURL.GET("validate-token", dependencies.Handlers.UserHandler.ValidateToken)
		}
	}

	r.Run()
}
