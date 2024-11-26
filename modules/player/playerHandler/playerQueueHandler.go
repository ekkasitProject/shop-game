package playerHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface {
	}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}
