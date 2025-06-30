package repository

import "github.com/ArtemHvozdov/tg-game-bot_v2/internal/database"

// PlayerRepository handles database operations for players
type PlayerRepository struct {
	db *database.Database
}

// NewPlayerRepository creates a new player repository
func NewPlayerRepository(db *database.Database) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}