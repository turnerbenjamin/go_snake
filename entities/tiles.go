package entities

import (
	"fmt"

	"github.com/turnerbenjamin/go_snake/entities/tileCodes"
)

type tileDescription struct {
	ColorCode byte
	Chars     string
}

func (td *tileDescription) GetColour() byte{
	return td.ColorCode
}

func (td *tileDescription) GetChars() string{
	return td.Chars
}

var tileMap = map[byte]tileDescription{
	tileCodes.GrassA: {ColorCode: 77, Chars: "  "},
	tileCodes.GrassB: {ColorCode: 76, Chars: "  "},
	tileCodes.SnakeHead: {ColorCode: 196, Chars: "00"},
	tileCodes.SnakeHeadBlinking: {ColorCode: 196, Chars: "--"},
	tileCodes.SnakeHeadDead: {ColorCode: 196, Chars: "XX"},
	tileCodes.SnakeBodyA: {ColorCode: 198, Chars: "  "},
	tileCodes.SnakeBodyB: {ColorCode: 200, Chars: "  "},
	tileCodes.SnakeBodyC: {ColorCode: 202, Chars: "  "},
	tileCodes.SnakeTail: {ColorCode: 16, Chars: "  "},
	tileCodes.SnakeHeadEating: {ColorCode: 196, Chars: "**"},
	tileCodes.SnakeBodyDigestingA: {ColorCode: 198, Chars: "()"},
	tileCodes.SnakeBodyDigestingB: {ColorCode: 200, Chars: "()"},
	tileCodes.SnakeBodyDigestingC: {ColorCode: 202, Chars: "()"},
	tileCodes.SnakeTailEating: {ColorCode: 16, Chars: "()"},
	tileCodes.AppleA: {ColorCode: 77, Chars: "🍎"},
	tileCodes.AppleB: {ColorCode: 76, Chars: "🍎"},
}

func GetTileDescription(k byte) string {
	if d, exists := tileMap[k]; exists{
	 return fmt.Sprintf("\033[48;5;%dm%s\033[0m", d.ColorCode, d.Chars)
	}
	panic("Tile description not defined");
}
