package server

import (
	"log"

	"github.com/ekkasitProject/shop-game/modules/item/itemHandler"
	itemPb "github.com/ekkasitProject/shop-game/modules/item/itemPb"
	"github.com/ekkasitProject/shop-game/modules/item/itemRepository"
	"github.com/ekkasitProject/shop-game/modules/item/itemUsecase"
	"github.com/ekkasitProject/shop-game/pkg/grpccon"
)

func (s *server) itemService() {
	itemRepository := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository)
	itemHttpHandler := itemHandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemQueueHandler := itemHandler.NewItemGrpcHandler(itemUsecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, itemQueueHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()
	_ = itemHttpHandler
	_ = itemQueueHandler

	item := s.app.Group("/item_v1")

	// Health check
	item.GET("/", s.healthCheckService)
}
