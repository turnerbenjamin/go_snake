package entities

import (
	"math/rand"

	"github.com/turnerbenjamin/go_snake/entities/tileCodes"
	"github.com/turnerbenjamin/go_snake/utilities"
	"github.com/turnerbenjamin/go_snake/utilities/directions"
)

type snakeSegment struct{
	next *snakeSegment
	position utilities.Position
	baseTileCode byte
	parent *snake
	isHead bool
	isDigesting bool
}


func (ss *snakeSegment) getPiece() Piece {
	return ss.parent
}

func (ss *snakeSegment) getTileCode() byte{
	btc := ss.baseTileCode
	if ss == ss.parent.head && ss.parent.isDead{
		return tileCodes.SnakeHeadDead
	}

	// if ss == ss.parent.tail && ss != ss.parent.head{
		// btc = tileCodes.SnakeTail 
	// }
	if ss.isDigesting {
		return ss.getDigestingTileCodeVariant(btc)
	}
	return btc
}

func (ss *snakeSegment) getDigestingTileCodeVariant(btc byte) byte{
	switch (btc) {
	case tileCodes.SnakeHead:
		return tileCodes.SnakeHeadEating
	case tileCodes.SnakeHeadBlinking:
		return tileCodes.SnakeHeadEating
	case tileCodes.SnakeBodyA:
		return tileCodes.SnakeBodyDigestingA
	case tileCodes.SnakeBodyB:
		return tileCodes.SnakeBodyDigestingB
	case tileCodes.SnakeBodyC:
		return tileCodes.SnakeBodyDigestingC
	case tileCodes.SnakeTail:
		return tileCodes.SnakeTailEating
	default:
		panic("Invalid base tile code")
	}
}

type snake struct{
    level *level
	direction directions.Direction
	head *snakeSegment
	tail *snakeSegment
	length int
	isDead bool
}


func CreateSnake(l *level, pos utilities.Position, dir directions.Direction) *snake{

	s := &snake{
		level: l,
		direction: dir,
		length: 1,
	}

	s.head = &snakeSegment{
		position: pos, 
		baseTileCode: tileCodes.SnakeHead, 
		parent: s,
		isHead: true,
	}
	s.tail = s.head

	s.level.placePiece(s.head.position, s.head)

	return s
}

func (s *snake) Update(){
		s.move();
		s.updateSnakeHeadTileCode()
}

func (s *snake) GobbleGobble(){
	s.head.isDigesting = true
}

func (s *snake) UpdateDirection(direction directions.Direction){
	moveOnX := direction == directions.Left || direction == directions.Right
	moveOnY := direction == directions.Up || direction == directions.Down;

	if moveOnX && !(s.direction == directions.Left || s.direction == directions.Right){
		s.direction = direction
	}
	if moveOnY && !(s.direction == directions.Up || s.direction == directions.Down){
		s.direction = direction
	}
}

func (s *snake) updateSnakeHeadTileCode(){
    isBlinking := rand.Intn(100) < 5
	if isBlinking {
		s.head.baseTileCode = tileCodes.SnakeHeadBlinking
	}else{
		s.head.baseTileCode = tileCodes.SnakeHead
	}
}

func (s *snake) move(){
	si := s.tail
	if s.tail.isDigesting{
		s.tail.isDigesting = false
		s.growUp()
	}

	for !si.isHead{
		next := si.next
		s.level.placePiece(next.position, si)
		si.position = next.position
		if(next.isDigesting){
			next.isDigesting = false
			si.isDigesting = true
		}
		si = si.next
	}
	s.head.position = s.getHeadDestination()
    s.level.placePiece(s.head.position, s.head)
}

func (s *snake) growUp(){
	nt := &snakeSegment{
		next: s.tail,
		position: s.tail.position,
		parent: s.tail.parent,
		baseTileCode: s.getBaseTileCodeForBody(),
	}
	s.length++
	s.tail = nt
}

func (s *snake) dealWithDeath(){
	s.isDead = true
}

func (s *snake) getBaseTileCodeForBody() byte{
	segmentCount := (s.length - 1) / 3;

	switch {
	case segmentCount % 3 == 0:
		return tileCodes.SnakeBodyA
	case segmentCount % 3 == 1:
		return tileCodes.SnakeBodyB
	default:
		return tileCodes.SnakeBodyC
	}
}

func (s *snake) getHeadDestination() utilities.Position{
	x := s.head.position.X
	y := s.head.position.Y

	switch s.direction{
	case directions.Up: {y -= 1}
	case directions.Right: {x = (x + 1) % s.level.width}
	case directions.Down: {y = (y + 1) % s.level.height}
	case directions.Left: {x -= 1}
	}

	if(x < 0){x = s.level.width - 1}
	if(y < 0){y = s.level.height - 1}

	return utilities.Position{X: x, Y: y}
}