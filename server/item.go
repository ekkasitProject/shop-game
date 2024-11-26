package server

import (
	itemUsecase "github.com/ekkasitProject/shop-game/modules/item/ItemUsecase"
	"github.com/ekkasitProject/shop-game/modules/item/itemHandler"
	"github.com/ekkasitProject/shop-game/modules/item/itemRepository"
)

func (s *server) itemService() {
	itemRepository := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository)
	itemHttpHandler := itemHandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemQueueHandler := itemHandler.NewItemGrpcHandler(itemUsecase)

	_ = itemHttpHandler
	_ = itemQueueHandler

	item := s.app.Group("/item_v1")

	// Health check
	item.GET("/health", s.healthCheckService)
}
