package service

import (
	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/repository"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

// NotificationService handles notifications to players
type NotificationService struct {
	playerRepo *repository.PlayerRepository
	gameRepo   *repository.GameRepository
	bot        *telebot.Bot
	logger     *logrus.Logger
}

// NewNotificationService creates a new notification service
func NewNotificationService(
	playerRepo *repository.PlayerRepository,
	gameRepo *repository.GameRepository,
	bot *telebot.Bot,
	logger *logrus.Logger,
) *NotificationService {
	return &NotificationService{
		playerRepo: playerRepo,
		gameRepo:   gameRepo,
		bot:        bot,
		logger:     logger,
	}
}