package main

import (
	"context"
	"log"
	"os"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/pkg/database/migration"
)

func main() {
	ctx := context.Background()

	// Inintialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())
	switch cfg.App.Name {
	case "auth":
		migration.AuthMigrate(ctx, &cfg)
	case "player":
		migration.PlayerMigrate(ctx, &cfg)
	case "item":
	case "inventory":

	}
}
