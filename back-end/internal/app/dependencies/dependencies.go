package dependencies

import (
	"github.com/luizmarinhojr/metrics/internal/app/repository"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
	"github.com/luizmarinhojr/metrics/internal/http/api/handler"
	"gorm.io/gorm"
)

type Dependencies struct {
	BrokerHandler *handler.BrokerHandler
	MetricHandler *handler.MetricHandler
}

func Inject(db *gorm.DB) *Dependencies {
	repositories := repository.NewRepository(db)

	useCases := usecase.NewUseCase(repositories)

	handlers := handler.NewHandler(useCases)

	return &Dependencies{
		BrokerHandler: handlers.BrokerHandler,
		MetricHandler: handlers.MetricHandler,
	}
}
