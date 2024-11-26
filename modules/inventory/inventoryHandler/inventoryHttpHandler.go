package inventoryHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface {
	}

	InventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &InventoryHttpHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}
