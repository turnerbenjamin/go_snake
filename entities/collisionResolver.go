package entities

import (
	"github.com/turnerbenjamin/go_snake/entities/collisionTypes"
)


type collisionResolver struct {
	handleSnakeAppleCollision func()
	handleSnakeSnakeCollision func()
}

func createCollisionResolver(
	onSnakeAppleCollision func(), 
	onSnakeSnakeCollision func()) *collisionResolver {
	return &collisionResolver{
		handleSnakeAppleCollision: onSnakeAppleCollision,
		handleSnakeSnakeCollision: onSnakeSnakeCollision,
	}
}

func (cr *collisionResolver) resolve(pu1 PieceUnit, pu2 PieceUnit) PieceUnit {

	snakePieceUnit, otherPieceUnit := cr.ensureSnakeFirstPieceUnit(pu1, pu2)
	switch cr.detectCollisionType(otherPieceUnit){
	case collisionTypes.SNAKE_APPLE:
		cr.handleSnakeAppleCollision()
	case collisionTypes.SNAKE_SNAKE:
		cr.handleSnakeSnakeCollision()
	default:
		panic("Invalid collision detected")
	}
	return snakePieceUnit
}

func (cr *collisionResolver) detectCollisionType(
	otherPieceUnit PieceUnit) byte {

    switch {
    case isSnake(otherPieceUnit.getPiece()):
        return collisionTypes.SNAKE_SNAKE
    case isApple(otherPieceUnit.getPiece()):
        return collisionTypes.SNAKE_APPLE
    default:
		return collisionTypes.INVALID
	}
}

func (cr *collisionResolver) ensureSnakeFirstPieceUnit(pu1, pu2 PieceUnit) (PieceUnit, PieceUnit){
	switch {
	case isSnake(pu1.getPiece()):
		return pu1, pu2
	case isSnake(pu2.getPiece()):
		return pu2, pu1
	default:
		panic("Collision must involve a snake")
	}
}

func isSnake(p Piece) bool {
    _, ok := p.(*snake)
    return ok
}

func isApple(p Piece) bool {
    _, ok := p.(*apple)
    return ok
}

