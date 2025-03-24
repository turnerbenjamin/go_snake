package interfaces

type Component interface {
	GetHeight() int
	GetWidth() int
	GetData() [][]byte
	GetTileDescriptionProvider() TileDescriptionProvider
}
