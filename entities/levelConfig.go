package entities

import (
	"github.com/turnerbenjamin/go_snake/utilities"
	"github.com/turnerbenjamin/go_snake/utilities/directions"
)

type LevelConfig struct {
	Width                 int
	Height                int
	SnakeStartingPos utilities.Position
	SnakeStartingDir directions.Direction
}