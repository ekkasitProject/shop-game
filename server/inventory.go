package server

import (
	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryHandler"
	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryRepository"
	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	inventoryRepository := inventoryRepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryUsecase.NewInventoryUsecase(inventoryRepository)
	inventoryHttpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryQueueHandler := inventoryHandler.NewInventoryGrpcHandler(inventoryUsecase)

	_ = inventoryHttpHandler
	_ = inventoryQueueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health check
	inventory.GET("/", s.healthCheckService)
}
