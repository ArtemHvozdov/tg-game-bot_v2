package models

type Player struct {
	ID       int64
	UserName string
	Name     string
	Status   string
	Skipped  int
	GameID   int
	Role     string // "admin", "player"
}

type GamePlayer struct {
	GameID   int
	PlayerID int
	Status   string // "joined", "playing", "finished"
}

type PlayerResponse struct {
	ID          int
	PlayerID    int64
	GameID      int
	TaskID      int
	HasResponse bool
	Skipped     bool
}

type SkipStatus struct {
	AlreadyAnswered  bool // true, if player already answered
	AlreadySkipped   bool // true, if player already skipped this task
	SkipLimitReached bool // true, if player already has three skips
	RemainingSkips   int  // number of remaining skips
}

type AddResponseResult struct {
	AlreadyAnswered bool
	AlreadySkipped  bool
	Success         bool
}
