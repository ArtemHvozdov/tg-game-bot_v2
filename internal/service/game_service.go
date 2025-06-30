package service

import (
	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/repository"
	"github.com/sirupsen/logrus"
)

// GameService provides game-related functionality
type GameService struct {
	gameRepo   *repository.GameRepository
	playerRepo *repository.PlayerRepository
	taskRepo   *repository.TaskRepository
	logger     *logrus.Logger
}

// NewGameService creates a new game service
func NewGameService(
	gameRepo *repository.GameRepository,
	playerRepo *repository.PlayerRepository,
	taskRepo *repository.TaskRepository,
	logger *logrus.Logger,
) *GameService {
	return &GameService{
		gameRepo:   gameRepo,
		playerRepo: playerRepo, 
		taskRepo:   taskRepo,
		logger:     logger,
	}
}
