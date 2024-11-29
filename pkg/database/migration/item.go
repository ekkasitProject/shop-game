package migration

import (
	"context"
	"log"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/item"
	"github.com/ekkasitProject/shop-game/pkg/database"
	"github.com/ekkasitProject/shop-game/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func itemDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("item_db")
}

func ItemMigrate(pctx context.Context, cfg *config.Config) {
	db := itemDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	collection := db.Collection("items")

	indexs, _ := collection.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "title", Value: 1}}},
	})
	for _, index := range indexs {
		log.Println("Index", index)
	}
	// roles data
	document := func() []any {
		roles := []*item.Item{
			{
				Title:       "Diamond Sword",
				Price:       1000,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      100,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Iron Sword",
				Price:       500,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      50,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Wooden Sword",
				Price:       100,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      20,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()
	// insert many
	result, err := collection.InsertMany(pctx, document)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrate item complete :", result.InsertedIDs)
}
