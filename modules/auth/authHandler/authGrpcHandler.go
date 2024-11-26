package authHandler

import (
	authUsecase "github.com/ekkasitProject/shop-game/modules/auth/authUsecase"
)

type (
	authGrpcHandlerService struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandlerService(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandlerService {
	return &authGrpcHandlerService{
		authUsecase: authUsecase,
	}
}
