package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/config"
	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/database"
	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/repository"
	"github.com/ArtemHvozdov/tg-game-bot_v2/internal/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

// Bot represents a Telegram bot instance with all its dependencies
type Bot struct {
	bot             *telebot.Bot
	gameService     *service.GameService
	taskService     *service.TaskService
	notifService    *service.NotificationService
	gameRepository  *repository.GameRepository
	playerRepository *repository.PlayerRepository
	taskRepository  *repository.TaskRepository
	logger          *logrus.Logger
	config          *config.Config
}

func NewBot(cfg *config.Config, db *database.Database, logger *logrus.Logger) (*Bot, error) {
	// Configure the bot settings
	settings := telebot.Settings{
		Token:  cfg.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	// Create the bot
	b, err := telebot.NewBot(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	// Initialize repositories
	gameRepo := repository.NewGameRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Initialize services
	gameService := service.NewGameService(gameRepo, playerRepo, taskRepo, logger)
	taskService := service.NewTaskService(taskRepo, gameRepo, playerRepo, logger)
	notifService := service.NewNotificationService(playerRepo, gameRepo, b, logger)

	bot := &Bot{
		bot:             b,
		gameService:     gameService,
		taskService:     taskService,
		notifService:    notifService,
		gameRepository:  gameRepo,
		playerRepository: playerRepo,
		taskRepository:  taskRepo,
		logger:          logger,
		config:          cfg,
	}
	
	// Register handlers
	bot.registerHandlers()

	return bot, nil
}

// Start starts the bot
func (b *Bot) Start() {
	b.bot.Start()
}

// Stop gracefully stops the bot
func (b *Bot) Stop(ctx context.Context) error {
	b.bot.Stop()
	return nil
}

func (b *Bot) registerHandlers() {
	// Command handlers
	b.bot.Handle("/start", b.StartHandler())
	b.bot.Handle("/help", b.HelpMeHandler())

	// Button handlers
	b.bot.Handle(&telebot.InlineButton{Data: "help_me"}, b.HelpMeHandler())

}