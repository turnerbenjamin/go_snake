package interfaces

import "github.com/turnerbenjamin/go_snake/utilities/directions"

type Ui interface {
	RenderWelcomeScreen()
	Init()
	RenderComponent(Component, int)
	ShowGameOverMessage(int) bool
	CleanUp()
	CheckForUserInput() (bool, directions.Direction, string)
}
