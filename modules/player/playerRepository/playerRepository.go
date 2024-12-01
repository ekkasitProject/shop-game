package playerRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ekkasitProject/shop-game/modules/player"
	"github.com/ekkasitProject/shop-game/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PlayerRepositoryService interface {
		IsUniquePlayer(pctx context.Context, email, username string) bool
		InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error)
		InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error
		GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error)
	}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepositoryService {
	return &playerRepository{
		db: db,
	}
}

func (r *playerRepository) playerDbConn() *mongo.Database {
	return r.db.Database("player_db")
}
func (r *playerRepository) IsUniquePlayer(pctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn()
	col := db.Collection("players")

	player := new(player.Player)
	if err := col.FindOne(
		ctx,
		bson.M{"%or": []bson.M{
			{"username": username},
			{"email": email},
		}},
	).Decode(player); err != nil {
		return true
	}
	return false
}
func (r *playerRepository) InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn()
	collection := db.Collection("players")
	playerId, err := collection.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayer: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one player failed")
	}
	return playerId.InsertedID.(primitive.ObjectID), nil
}

func (r *playerRepository) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn()
	collection := db.Collection("players")

	result := new(player.PlayerProfileBson)

	if err := collection.FindOne(
		ctx,
		bson.M{"_id": utils.ConventToObjectId(playerId)},
		options.FindOne().SetProjection(bson.M{
			"id":         1,
			"email":      1,
			"username":   1,
			"created_at": 1,
			"updated_at": 1,
		}),
	).Decode(result); err != nil {
		log.Printf("Error: FindOnePlayerProfile: %s", err.Error())
		return nil, errors.New("error: find one player profile failed")
	}

	return result, nil
}

func (r *playerRepository) InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn()
	collection := db.Collection("player_transactions")

	result, err := collection.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayerTransaction: %s", err.Error())
		return errors.New("error: Insert One Player Transaction failed")
	}
	log.Printf("Result: InsertOnePlayerTransaction: %v", result.InsertedID)
	return nil
}

func (r *playerRepository) GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn()
	collection := db.Collection("player_transactions")

	filter := bson.A{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "player_id", Value: playerId},
			}},
		},
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$player_id"},
				{Key: "balance", Value: bson.D{
					{Key: "$sum", Value: "$amount"},
				}},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "player_id", Value: "$_id"},
				{Key: "_id", Value: 0},
				{Key: "balance", Value: 1},
			}},
		},
	}

	cursors, err := collection.Aggregate(ctx, filter)
	if err != nil {
		log.Printf("Error: GetPlayerSavingAccount: %s", err.Error())
		return nil, errors.New("error: Get Player Saving Account failed")
	}

	result := new(player.PlayerSavingAccount)

	for cursors.Next(ctx) {
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: GetPlayerSavingAccount: %s", err.Error())
			return nil, errors.New("error: Get Player Saving Account failed")
		}
	}
	return result, nil
}
