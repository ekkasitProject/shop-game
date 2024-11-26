package inventoryHandler

import "github.com/ekkasitProject/shop-game/modules/inventory/inventoryUsecase"

type (
	inventoryGrpcHandlerService struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandlerService {
	return &inventoryGrpcHandlerService{
		inventoryUsecase: inventoryUsecase,
	}
}
