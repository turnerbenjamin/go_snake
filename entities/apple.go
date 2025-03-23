package entities

import (
	"github.com/turnerbenjamin/go_snake/entities/tileCodes"
	"github.com/turnerbenjamin/go_snake/utilities"
)

type apple struct {
	level    *level
	position *utilities.Position
}

func createApple(l *level) *apple{
	a := &apple{
		level: l,
		position: l.getRandomAvailableSpace(),
	}
	
	a.level.placePiece(*a.position, a)

	return a
}


func (a *apple) getPiece() Piece {
	return a
}

func (a *apple) getTileCode() byte {
	if (a.position.X + a.position.Y) % 2 == 0 {
		return tileCodes.AppleA
	}
	return tileCodes.AppleB
}

func (a *apple) Update() {
	a.level.placePiece(*a.position, a)
}

func (a *apple) randomisePosition() {
	a.position = a.level.getRandomAvailableSpace();
}

