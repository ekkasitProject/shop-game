package inventoryUsecase

import "github.com/ekkasitProject/shop-game/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface {
	}
	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}
