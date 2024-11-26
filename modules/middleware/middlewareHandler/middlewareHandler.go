package middlewareHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface {
	}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}
