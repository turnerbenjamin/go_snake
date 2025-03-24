package game

import "github.com/turnerbenjamin/go_snake/interfaces"

type GameConfig struct {
	Level          interfaces.Level
	Ui             interfaces.Ui
	RateLimit      int
	TurnDurationMs int
	FpsSampleSize  int
}
