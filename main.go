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
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	// Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	// Start Server
	server.Start(ctx, &cfg, db)
}
