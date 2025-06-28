#!/bin/bash

# Ensure the script is executed inside "friend-game-bot"
if [[ $(basename "$PWD") != "tg-game-bot_v2" ]]; then
    echo "Error: Please run this script inside the 'friend-game-bot' directory."
    exit 1
fi

# Define all directories to be created
DIRECTORIES=(
    "cmd/bot"
    "internal/bot"
    "internal/config"
    "internal/database/migrations"
    "internal/game"
    "internal/models"
    "internal/repository"
    "internal/service"
    "pkg/cache"
    "pkg/utils"
    "assets/tasks"
)

# Define all files to be created
FILES=(
    "cmd/bot/main.go"
    "internal/bot/bot.go"
    "internal/bot/handlers.go"
    "internal/bot/middleware.go"
    "internal/config/config.go"
    "internal/database/database.go"
    "internal/database/migrations/schema.sql"
    "internal/game/game.go"
    "internal/game/player.go"
    "internal/game/task.go"
    "internal/models/game.go"
    "internal/models/player.go"
    "internal/models/task.go"
    "internal/repository/game_repository.go"
    "internal/repository/player_repository.go"
    "internal/repository/task_repository.go"
    "internal/service/game_service.go"
    "internal/service/notification_service.go"
    "internal/service/task_service.go"
    "pkg/cache/redis.go"
    "pkg/utils/link_generator.go"
    "pkg/utils/logger.go"
    "assets/tasks/task_definitions.json"
    ".env"
    ".gitignore"
    "docker-compose.yml"
    "Dockerfile"
    "go.mod"
    "README.md"
)

# Create directories
echo "Creating directories..."
for dir in "${DIRECTORIES[@]}"; do
    mkdir -p "$dir"
    echo "Created: $dir"
done

# Create files
echo "Creating files..."
for file in "${FILES[@]}"; do
    touch "$file"
    echo "Created: $file"
done

echo "✅ File structure successfully generated inside 'friend-game-bot'."
