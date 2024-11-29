package migration

import (
	"context"
	"log"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/auth"
	"github.com/ekkasitProject/shop-game/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func authDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("auth_db")
}

func AuthMigrate(pctx context.Context, cfg *config.Config) {
	db := authDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	collection := db.Collection("auth")

	// indexs
	indexs, _ := collection.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "player_id", Value: 1}}},
		{Keys: bson.D{{Key: "refresh_token", Value: 1}}},
	})
	for _, index := range indexs {
		log.Println("Index", index)
	}

	// roles
	collection = db.Collection("roles")
	// indexs
	indexs, _ = collection.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "code", Value: 1}}},
	})
	for _, index := range indexs {
		log.Println("Index", index)
	}
	// roles data
	document := func() []any {
		roles := []*auth.Role{
			{Title: "player",
				Code: 0,
			},
			{Title: "admin",
				Code: 1,
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
	log.Println("Migrate auth complete :", result.InsertedIDs)
}
