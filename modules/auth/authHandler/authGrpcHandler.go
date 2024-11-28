package authHandler

import (
	"context"

	authPb "github.com/ekkasitProject/shop-game/modules/auth/authPb"
	authUsecase "github.com/ekkasitProject/shop-game/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
		authPb.UnimplementedAuthGrpcServiceServer
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{
		authUsecase: authUsecase,
	}
}

func (g *authGrpcHandler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return &authPb.AccessTokenSearchRes{}, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return &authPb.RolesCountRes{}, nil
}
