package interfaces

type TileDescriptionProvider interface {
	GetTileDescription(byte) TileDescription
}

type TileDescription interface {
	GetBackgroundColour() byte
	GetText() string
}
