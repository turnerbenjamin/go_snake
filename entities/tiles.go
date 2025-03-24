package entities

import (
	"github.com/turnerbenjamin/go_snake/entities/tileCodes"
	"github.com/turnerbenjamin/go_snake/interfaces"
)

type tileDescription struct {
	ColorCode byte
	Chars     string
}

func (td *tileDescription) GetBackgroundColour() byte {
	return td.ColorCode
}

func (td *tileDescription) GetText() string {
	return td.Chars
}

type snakeTileDescriptionProvider struct {
	tileMap map[byte]*tileDescription
}

func (p *snakeTileDescriptionProvider) GetTileDescription(k byte) interfaces.TileDescription {
	if d, exists := p.tileMap[k]; exists {
		return d
	}
	panic("Tile description not defined")
}

func newSnakeTileDescriptionProvider() *snakeTileDescriptionProvider {
	return &snakeTileDescriptionProvider{
		tileMap: map[byte]*tileDescription{
			tileCodes.GrassA:              {ColorCode: 77, Chars: "  "},
			tileCodes.GrassB:              {ColorCode: 76, Chars: "  "},
			tileCodes.SnakeHead:           {ColorCode: 196, Chars: "00"},
			tileCodes.SnakeHeadBlinking:   {ColorCode: 196, Chars: "--"},
			tileCodes.SnakeHeadDead:       {ColorCode: 196, Chars: "XX"},
			tileCodes.SnakeBodyA:          {ColorCode: 198, Chars: "  "},
			tileCodes.SnakeBodyB:          {ColorCode: 200, Chars: "  "},
			tileCodes.SnakeBodyC:          {ColorCode: 202, Chars: "  "},
			tileCodes.SnakeTail:           {ColorCode: 16, Chars: "  "},
			tileCodes.SnakeHeadEating:     {ColorCode: 196, Chars: "**"},
			tileCodes.SnakeBodyDigestingA: {ColorCode: 198, Chars: "()"},
			tileCodes.SnakeBodyDigestingB: {ColorCode: 200, Chars: "()"},
			tileCodes.SnakeBodyDigestingC: {ColorCode: 202, Chars: "()"},
			tileCodes.SnakeTailEating:     {ColorCode: 16, Chars: "()"},
			tileCodes.AppleA:              {ColorCode: 77, Chars: "🍎"},
			tileCodes.AppleB:              {ColorCode: 76, Chars: "🍎"},
		},
	}
}
