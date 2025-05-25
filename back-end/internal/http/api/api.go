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
			baseURL.GET("/brokers", dependencies.BrokerHandler.GetAll)
			baseURL.GET("/broker/id", dependencies.BrokerHandler.GetById)
			baseURL.GET("/broker/name", dependencies.BrokerHandler.GetByName)
			baseURL.POST("/broker", dependencies.BrokerHandler.Create)
		}
		{
			baseURL.GET("metrics", dependencies.MetricHandler.GetAll)
			baseURL.GET("metrics/id")
			baseURL.POST("metric")
		}
	}

	r.Run()
}
