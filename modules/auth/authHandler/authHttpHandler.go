package authHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	authUsecase "github.com/ekkasitProject/shop-game/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		cfg:         cfg,
		authUsecase: authUsecase,
	}
}
