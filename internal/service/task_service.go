package service

import (
	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/repository"
	"github.com/sirupsen/logrus"
)

// TaskService provides task-related functionality
type TaskService struct {
	taskRepo   *repository.TaskRepository
	gameRepo   *repository.GameRepository
	playerRepo *repository.PlayerRepository
	logger     *logrus.Logger
}

// NewTaskService creates a new task service
func NewTaskService(
	taskRepo *repository.TaskRepository,
	gameRepo *repository.GameRepository,
	playerRepo *repository.PlayerRepository,
	logger *logrus.Logger,
) *TaskService {
	return &TaskService{
		taskRepo:   taskRepo,
		gameRepo:   gameRepo,
		playerRepo: playerRepo,
		logger:     logger,
	}
}