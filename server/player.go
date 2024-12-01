package server

import (
	"log"

	"github.com/ekkasitProject/shop-game/modules/player/playerHandler"
	playerPb "github.com/ekkasitProject/shop-game/modules/player/playerPb"
	"github.com/ekkasitProject/shop-game/modules/player/playerRepository"
	"github.com/ekkasitProject/shop-game/modules/player/playerUsecase"
	"github.com/ekkasitProject/shop-game/pkg/grpccon"
)

func (s *server) playerService() {
	authRepository := playerRepository.NewPlayerRepository(s.db)
	authUsecase := playerUsecase.NewPlayerUsecase(authRepository)
	playerHttpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, authUsecase)
	playerGrpcHandler := playerHandler.NewPlayerGrpcHandler(authUsecase)
	playerQueueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, authUsecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, playerGrpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = playerQueueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("/", s.healthCheckService)
	player.POST("/register", playerHttpHandler.CreatePlayer)
	player.POST("/add-money/", playerHttpHandler.AddPlayerMoney, s.middleware.JwtAuthorization)
	player.GET("/:player_id", playerHttpHandler.FindOnePlayerProfile)
	player.GET("/saving-account", playerHttpHandler.GetPlayerSavingAccount, s.middleware.JwtAuthorization)
}
