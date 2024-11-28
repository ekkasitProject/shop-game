package inventoryHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface {
	}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}
