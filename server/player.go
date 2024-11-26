package server

import (
	"github.com/ekkasitProject/shop-game/modules/player/playerHandler"
	"github.com/ekkasitProject/shop-game/modules/player/playerRepository"
	"github.com/ekkasitProject/shop-game/modules/player/playerUsecase"
)

func (s *server) playerService() {
	authRepository := playerRepository.NewPlayerRepository(s.db)
	authUsecase := playerUsecase.NewPlayerUsecase(authRepository)
	playerHttpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, authUsecase)
	playerGrpcHandler := playerHandler.NewPlayerGrpcHandler(authUsecase)
	playerQueueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, authUsecase)
	_ = playerHttpHandler
	_ = playerGrpcHandler
	_ = playerQueueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("/health", s.healthCheckService)
}
