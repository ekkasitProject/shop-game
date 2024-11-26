package itemHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface {
	}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		cfg:         cfg,
		itemUsecase: itemUsecase,
	}
}
