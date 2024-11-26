package main

import (
	"context"
	"log"
	"os"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/pkg/database"
	"github.com/ekkasitProject/shop-game/server"
)

func main() {
	ctx := context.Background()
	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Please provide config path")
		}

		return os.Args[1]
	}())

	// Initialize database
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	// Start server
	server.Start(ctx, &cfg, db)
}
