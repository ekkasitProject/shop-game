package database

import (
	"context"
	"log"
	"time"

	"github.com/ekkasitProject/shop-game/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return client
}
