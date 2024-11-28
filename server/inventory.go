package server

import (
	"log"

	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryHandler"
	inventoryPb "github.com/ekkasitProject/shop-game/modules/inventory/inventoryPb"
	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryRepository"
	"github.com/ekkasitProject/shop-game/modules/inventory/inventoryUsecase"
	"github.com/ekkasitProject/shop-game/pkg/grpccon"
)

func (s *server) inventoryService() {
	inventoryRepository := inventoryRepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryUsecase.NewInventoryUsecase(inventoryRepository)
	inventoryHttpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpcHandler := inventoryHandler.NewInventoryGrpcHandler(inventoryUsecase)
	inventoryQueueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, inventoryUsecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, inventoryGrpcHandler)

		log.Printf("Inventory gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()
	_ = inventoryHttpHandler
	_ = inventoryQueueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health check
	inventory.GET("/", s.healthCheckService)
}
