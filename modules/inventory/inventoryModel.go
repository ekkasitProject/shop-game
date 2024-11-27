package inventory

import (
	"github.com/ekkasitProject/shop-game/modules/item"
	"github.com/ekkasitProject/shop-game/modules/models"
)

type (
	UpdateInventoryReq struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}
	ItemInInventory struct {
		InventoryId string `json:"inventory_id"`
		*item.ItemShowCase
	}
	PlayerInventory struct {
		PlayerId string `json:"player_id"`
		*models.PaginteRes
	}
)
