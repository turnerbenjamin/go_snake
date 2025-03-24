package interfaces

import "github.com/turnerbenjamin/go_snake/utilities/directions"

type Level interface {
	Component
	NewGame()
	IsRunning() bool
	GetApplesEaten() int
	Update()
	HandleDirectionInput(directions.Direction)
}
