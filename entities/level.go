package entities

import (
	"math/rand"

	"github.com/turnerbenjamin/go_snake/entities/tileCodes"
	"github.com/turnerbenjamin/go_snake/utilities"
	"github.com/turnerbenjamin/go_snake/utilities/directions"
)

type level struct {
	width, height int
	data          [][]byte
	isRunning bool
	applesEaten int
	availableSpaces utilities.List[int]
	piecePlacements map[int]PieceUnit
	collisionResolver *collisionResolver
	apple *apple
	snake *snake
	snakeStartingPos utilities.Position
	snakeStartingDir directions.Direction
}

func CreateLevel(c LevelConfig) *level {
	l := level{
		width: c.Width,
		height: c.Height,
		snakeStartingPos: c.SnakeStartingPos,
		snakeStartingDir: c.SnakeStartingDir,
		availableSpaces: utilities.NewList[int](c.Width*c.Height),
		piecePlacements: make(map[int]PieceUnit),

	}

	return &l
}

func (l *level) NewGame(){
	l.isRunning = true
	l.applesEaten = 0

	l.InitialisePieces(l.snakeStartingPos, l.snakeStartingDir)
	l.InitialiseCollisionResolver()
	l.buildData()
}

func (l *level) IsRunning() bool{
	return l.isRunning
}

func (l *level) GetApplesEaten() int{
	return l.applesEaten
}

func (l *level) GetHeight() int{
	return l.height
}

func (l *level) GetWidth() int{
	return l.width
}

func (l *level) GetData() [][]byte {
	return l.data
}



func (l *level) InitialisePieces(snakeStartingPos utilities.Position, snakeStartingDir directions.Direction){
    l.snake = CreateSnake(
		l,
		snakeStartingPos, 
		snakeStartingDir,	 
	)
	l.apple = createApple(l)
}

func (l *level) InitialiseCollisionResolver(){
	l.collisionResolver = createCollisionResolver(
		l.handleSnakeAppleCollision,
		l.handleSnakeSnakeCollision,
	)
}

func (l *level) Update(){
	l.piecePlacements = make(map[int]PieceUnit)
	l.snake.Update()
	l.apple.Update()
	l.buildData()
}

func (l *level) HandleDirectionInput(d directions.Direction){
	l.snake.UpdateDirection(d)
}

func (l *level) buildData(){
	data := make([][]byte, l.height)

	for y := range l.height {
		data[y] = make([]byte, l.width)
		for x := range l.width {
			cid := l.getCellId(y,x)
			piecePlacement, isPiecePlacement := l.piecePlacements[cid]
			if(isPiecePlacement){
				data[y][x] = piecePlacement.getTileCode()
			}else{
				data[y][x] = getGrassTileCode(y,x)
			}
		}
	}
	l.data = data
}


func (l *level) placePiece(pos utilities.Position, p PieceUnit) {
	cid := l.getCellId(pos.Y, pos.X)
	if op, isCollision := l.piecePlacements[cid]; isCollision {
		p = l.collisionResolver.resolve(p, op)
	}
	l.piecePlacements[cid] = p;
}




func (l *level) getCellId(y,x int) int{
	return y * l.width + x
}

func (l *level) parseCellId(id int) *utilities.Position{
	return &utilities.Position{
		X: id % l.width,
		Y: id / l.width,
	}
}


func (l *level) getRandomAvailableSpace() *utilities.Position{
	l.updateAvailableSpaces()
	spaceI := rand.Intn(l.availableSpaces.Size());
    cid := l.availableSpaces.Get(spaceI);
	return l.parseCellId(cid)
}

func (l *level) updateAvailableSpaces(){
	l.availableSpaces.Clear()
	for y := range l.height {
		for x := range l.width {
			cid := l.getCellId(y,x)
			_, isPiecePlacement := l.piecePlacements[cid]
			if(!isPiecePlacement){
				l.availableSpaces.Push(cid)
			}
		}
	}
}


func (l *level) handleSnakeAppleCollision(){
	l.applesEaten++

	l.apple.randomisePosition()
	l.apple.Update()

	l.snake.GobbleGobble()
}

func (l *level) handleSnakeSnakeCollision(){
	l.snake.dealWithDeath()
	l.isRunning = false
}

func getGrassTileCode(y,x int) byte{
	isEven := (y + x) % 2 == 0;
	if isEven {
		return tileCodes.GrassA
	}
	return tileCodes.GrassB
}

