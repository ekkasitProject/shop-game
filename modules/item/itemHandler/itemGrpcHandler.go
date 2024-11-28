package itemHandler

import (
	"context"

	itemPb "github.com/ekkasitProject/shop-game/modules/item/itemPb"
	"github.com/ekkasitProject/shop-game/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemUsecase.ItemUsecaseService
		itemPb.UnimplementedItemGrpcServiceServer
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}
func (g *itemGrpcHandler) FindItemInIds(ctx context.Context, req *itemPb.FindItemInIdsReq) (*itemPb.FindItemInIdsRes, error) {
	return &itemPb.FindItemInIdsRes{}, nil
}
