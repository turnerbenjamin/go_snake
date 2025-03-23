package game

type GameConfig struct {
	Level Level
	Ui Ui
	RateLimit   int
	TurnDurationMs int
	FpsSampleSize int
}