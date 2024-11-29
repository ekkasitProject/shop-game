package migration

import (
	"context"
	"log"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func inventoryDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("inventory_db")
}

func InventoryMigrate(pctx context.Context, cfg *config.Config) {
	db := inventoryDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	collection := db.Collection("players_inventory")

	indexs, _ := collection.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "player_id", Value: 1}, {Key: "item_id", Value: 1}}},
	})
	for _, index := range indexs {
		log.Println("Index", index)
	}

	collection = db.Collection("players_inventory_queue")
	// insert many
	result, err := collection.InsertOne(pctx, bson.M{"offset": -1})
	if err != nil {
		panic(err)
	}
	log.Println("Migrate Inventory complete :", result)
}
