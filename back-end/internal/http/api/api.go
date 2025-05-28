package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/internal/app/dependencies"
)

func InitializeApi(dependencies *dependencies.Dependencies) {
	r := gin.Default()

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
		}
	}

	r.Run()
}
