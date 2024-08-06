package main

import (
	"icon-pln/internal/app"
	"icon-pln/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server, cleanup, err := app.BootstrapApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	defer cleanup()

	if err := server.Start(cfg.Server.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
