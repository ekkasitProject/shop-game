package server

import (
	"github.com/ekkasitProject/shop-game/modules/auth/authHandler"
	"github.com/ekkasitProject/shop-game/modules/auth/authRepository"
	"github.com/ekkasitProject/shop-game/modules/auth/authUsecase"
)

func (s *server) authService() {
	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecase(authRepository)
	authHttpHandler := authHandler.NewAuthHttpHandler(s.cfg, authUsecase)
	authGrpcHandler := authHandler.NewAuthGrpcHandlerService(authUsecase)

	_ = authHttpHandler
	_ = authGrpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("/health", s.healthCheckService)
}
