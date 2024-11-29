package migration

import (
	"context"
	"log"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/player"
	"github.com/ekkasitProject/shop-game/pkg/database"
	"github.com/ekkasitProject/shop-game/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func playerDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("player_db")
}

func PlayerMigrate(pctx context.Context, cfg *config.Config) {
	db := playerDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	// player_transactions
	collection := db.Collection("player_transactions")

	// indexs
	indexs, _ := collection.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "player_id", Value: 1}}},
	})
	log.Println(indexs)
	// players
	collection = db.Collection("players")
	// indexs
	indexs, _ = collection.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "email", Value: 1}}},
	})
	log.Println(indexs)

	documents := func() []any {
		roles := []*player.Player{
			{
				Email:    "player1@gmail.com",
				Password: "123456",
				Username: "player001",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player2@gmail.com",
				Password: "123456",
				Username: "player002",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player3@gmail.com",
				Password: "123456",
				Username: "player003",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "admin1@gmail.com",
				Password: "123456",
				Username: "admin001",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}
		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := collection.InsertMany(pctx, documents)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrate auth complete :", results.InsertedIDs)

	// Create player transactions
	playerTransactions := make([]any, 0)
	for _, p := range results.InsertedIDs {
		playerTransactions = append(playerTransactions, &player.PlayerTransaction{
			PlayerId:  "player:" + p.(primitive.ObjectID).Hex(),
			Amount:    1000,
			CreatedAt: utils.LocalTime().Local(),
		})
	}
	// Insert the transactions into player_transactions collection
	transactionCollection := db.Collection("player_transactions")
	transactionResults, err := transactionCollection.InsertMany(pctx, playerTransactions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Player transactions migration complete:", transactionResults.InsertedIDs)

	// Fix the syntax errors in this part
	collection = db.Collection("player_transaction_queue")
	result, err := collection.InsertOne(pctx, bson.M{"offset": -1})
	if err != nil {
		panic(err)
	}
	log.Println("Migrate Player complete :", result)
}
