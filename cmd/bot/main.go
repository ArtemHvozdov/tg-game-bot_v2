package main

import (
	"log"

	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/config"
	"github.com/ArtemHvozdov/tg-game-bot_v2/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviroment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Initialize logger
	logger := utils.NewLogger()
	logger.Info("Starting Friend Game Bot")

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	dataDir := cfg.DatabaseDir
	dataFile := cfg.DatabaseFile

	dbUrl := dataDir + dataFile

	db, err := database.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
}