package migration

import (
	"context"
	"log"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func paymentDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("payment_db")
}

func PaymentMigrate(pctx context.Context, cfg *config.Config) {
	db := paymentDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	collection := db.Collection("payment_queue")
	// insert many
	result, err := collection.InsertOne(pctx, bson.M{"offset": -1})
	if err != nil {
		panic(err)
	}
	log.Println("Migrate payment complete :", result)
}
