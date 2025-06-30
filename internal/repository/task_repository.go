package repository

import "github.com/ArtemHvozdov/tg-game-bot_v2/internal/database"

// TaskRepository handles database operations for tasks
type TaskRepository struct {
	db *database.Database
}

// NewTaskRepository creates a new task repository
func NewTaskRepository(db *database.Database) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}