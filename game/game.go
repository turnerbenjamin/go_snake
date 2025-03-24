package game

import (
	"time"

	"github.com/turnerbenjamin/go_snake/interfaces"
	"github.com/turnerbenjamin/go_snake/utilities/directions"
)

type game struct {
	level             interfaces.Level
	ui                interfaces.Ui
	stats             *stats
	rateLimitDuration time.Duration
	framesInMove      int
	turnDurationMs    int
	score             int
	isRunning         bool
}

func Create(c GameConfig) *game {
	g := &game{
		level:             c.Level,
		ui:                c.Ui,
		turnDurationMs:    c.TurnDurationMs,
		rateLimitDuration: time.Second / time.Duration(c.RateLimit),
		stats:             newStats(c.FpsSampleSize, c.RateLimit),
	}
	return g
}

func (g *game) Start() {
	g.ui.Init()
	defer g.ui.CleanUp()
	g.isRunning = true
	g.ui.RenderWelcomeScreen()
	g.loop()
}

func (g *game) loop() {

	for g.isRunning {
		g.playLevel()
		if !g.isRunning {
			break
		}
		g.isRunning = g.ui.ShowGameOverMessage(g.score)
	}
}

func (g *game) playLevel() {
	g.score = 0
	g.level.NewGame()
	g.stats.startTracking()
	for g.isRunning && g.level.IsRunning() {
		g.update()
		g.render()
		g.stats.update()
		time.Sleep(g.rateLimitDuration)
		g.score = g.level.GetApplesEaten() * 10
	}
}

func (g *game) update() {
	framesPerMove := g.calculateFramesPerMove()
	g.framesInMove++

	inputEntered, dir, char := g.ui.CheckForUserInput()
	if inputEntered {
		if dir != directions.InvalidDirection {
			g.level.HandleDirectionInput(dir)
		} else {
			if char == "q" {
				g.isRunning = false
			}
		}
	}

	if (inputEntered && dir != directions.InvalidDirection) || g.framesInMove >= framesPerMove {
		g.level.Update()
		g.framesInMove = 0
	}
}

func (g *game) render() {
	g.ui.RenderComponent(g.level, g.score)
}

func (g *game) calculateFramesPerMove() int {
	return int(float64(g.stats.fps) / 1000 * float64(g.turnDurationMs))
}
