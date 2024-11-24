package main

import (
	"context"
	"log"
	"os"

	"github.com/ekkasitProject/shop-game/config"
)

func main() {
	ctx := context.Background()
	_ = ctx
	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Please provide config path")
		}

		return os.Args[1]
	}())
	log.Printf("Config: %+v", cfg)
}
