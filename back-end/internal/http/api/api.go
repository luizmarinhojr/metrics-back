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
			baseURL.GET("/broker/id")
			baseURL.GET("/broker/name")
		}
		{
			baseURL.GET("metrics")
			baseURL.GET("metrics/id")
		}
	}

	r.Run()
}
