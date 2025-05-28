package dependencies

import (
	"github.com/luizmarinhojr/metrics/internal/app/auth"
	"github.com/luizmarinhojr/metrics/internal/app/repository"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
	"github.com/luizmarinhojr/metrics/internal/http/api/handler"
	"github.com/luizmarinhojr/metrics/internal/http/api/middleware"
	"gorm.io/gorm"
)

type Dependencies struct {
	Handlers    *handler.Handler
	Middlewares *middleware.Middleware
}

func Inject(db *gorm.DB) *Dependencies {
	repositories := repository.NewRepository(db)
	useCases := usecase.NewUseCase(repositories)

	handlers := handler.NewHandler(useCases)
	middlewares := middleware.NewMiddleware(auth.NewAuth())

	return &Dependencies{
		Handlers:    handlers,
		Middlewares: middlewares,
	}
}
