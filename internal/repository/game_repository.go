package repository

import "github.com/ArtemHvozdov/tg-game-bot_v2/internal/database"

// GameRepository handles database operations for games
type GameRepository struct {
	db *database.Database
}

// NewGameRepository creates a new game repository
func NewGameRepository(db *database.Database) *GameRepository {
	return &GameRepository{
		db: db,
	}
}