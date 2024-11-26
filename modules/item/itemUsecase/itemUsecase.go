package itemUsecase

import "github.com/ekkasitProject/shop-game/modules/item/itemRepository"

type (
	ItemUsecaseService interface {
	}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{
		itemRepository: itemRepository,
	}
}
