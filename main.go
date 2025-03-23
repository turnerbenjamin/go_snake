package main

import (
	"github.com/turnerbenjamin/go_snake/entities"
	"github.com/turnerbenjamin/go_snake/game"
	"github.com/turnerbenjamin/go_snake/utilities"
	"github.com/turnerbenjamin/go_snake/utilities/directions"
	"github.com/turnerbenjamin/go_snake/view"
)


func main() {

	w := 20
	h := 20
    ssp := utilities.Position{X: w / 2, Y: h / 2}

	lc := entities.LevelConfig{
		Width: 20,
		Height: 20,
		SnakeStartingPos: ssp,
		SnakeStartingDir: directions.Right,
	}

    l := entities.CreateLevel(lc)

	ui := view.CreateUi()

	gc := game.GameConfig{
		Level: l,
		Ui: ui,
		RateLimit: 70,
		TurnDurationMs: 150,
		FpsSampleSize: 25,
	}

	g := game.Create(gc)
	g.Start()
}
