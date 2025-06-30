package models

type Game struct {
	ID            int
	Name          string
	GameChatID    int64
	CurrentTaskID int
	TotalPlayers  int    // max 5
	Status        string // "waiting", "playing", "finished"
}

type GameState struct {
	GameID  int64
	Current string // Current FSM state
}

// Game status constants
const (
	StatusGameWaiting  = "waiting"
	StatusGamePlaying  = "playing"
	StatusGameFinished = "finished"
)
