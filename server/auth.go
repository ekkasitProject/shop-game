package server

import (
	"log"

	"github.com/ekkasitProject/shop-game/modules/auth/authHandler"
	authPb "github.com/ekkasitProject/shop-game/modules/auth/authPb"
	"github.com/ekkasitProject/shop-game/modules/auth/authRepository"
	"github.com/ekkasitProject/shop-game/modules/auth/authUsecase"
	"github.com/ekkasitProject/shop-game/pkg/grpccon"
)

func (s *server) authService() {
	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecase(authRepository)
	authHttpHandler := authHandler.NewAuthHttpHandler(s.cfg, authUsecase)
	authGrpcHandler := authHandler.NewAuthGrpcHandler(authUsecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, authGrpcHandler)

		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = authHttpHandler
	_ = authGrpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("/", s.healthCheckService)
}
