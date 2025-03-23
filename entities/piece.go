package entities

type Piece interface {
	Update()
}

type PieceUnit interface {
	getPiece() Piece
	getTileCode() byte
}
