package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

)	

type Position struct {
	x int32
	y int32
}

type Player struct {
	dir Position
	body []Position
}

var initialBody = []Position{
	{8, 8}, {8,8}, {8,8},
}

var freeTiles []Position

var SquareSize int32
var GameSize int
var GameSpeed int
var FrameCount *int

var tileQty int32 = 16

var player = Player{
	body: append([]Position{}, initialBody...),
	dir: Position{x: 0, y: 0},
	
}

var apple = Position{5, 6}

func Start(){
	tileQty = int32(GameSize / int(SquareSize))

	setFreeTiles()
}

func playerDirection(){
	nextDir := player.dir

	if rl.IsKeyDown(rl.KeyW)  {
		nextDir.y = -1
		nextDir.x = 0
	}
	if rl.IsKeyDown(rl.KeyS) {
		nextDir.y = 1
		nextDir.x = 0
	}
	if rl.IsKeyDown(rl.KeyA) {
		nextDir.y = 0
		nextDir.x = -1
	}
	if rl.IsKeyDown(rl.KeyD) {
		nextDir.y = 0
		nextDir.x = 1
	}	

	if  (player.body[1] == Position{player.body[0].x + nextDir.x, player.body[0].y + nextDir.y} ) {
		return
	}
	
	player.dir = nextDir
}

func PlayerPosition() {

	if (player.dir == Position{x: 0, y: 0}){
		return
	}

	for i := len(player.body) -1; i >= 0 ; i-- {
		if i == 0 {
			player.body[i].x += player.dir.x
			player.body[i].y += player.dir.y

			if player.body[i].x >= tileQty{
				player.body[i].x = int32(0)
			}
	
			if player.body[i].y >= tileQty{
				player.body[i].y = int32(0)
			}
	
			if player.body[i].x < int32(0){
				player.body[i].x = tileQty
			}
	
			if player.body[i].y < int32(0){
				player.body[i].y = tileQty
			}

			continue
		}
		player.body[i].x = player.body[i - 1].x
		player.body[i].y = player.body[i - 1].y
	}
}

func Update() {

	playerDirection()

	if (*FrameCount % (60 / GameSpeed) == 0){
		PlayerPosition()
		collision()
	}

	drawApple()
	drawSnake()
}

func collision() {
	head:= player.body[0]

	if head.x == apple.x && head.y == apple.y{
		player.body = append(player.body, player.body[len(player.body) - 1] )
	}

	for i := 1; i < len(player.body); i++ {
		body := player.body[i]
		if body == head{
			setInitialConfig()
		}
	}
}

func setFreeTiles() {
	for x := 0; x < int(tileQty); x++ {
		for y := 0; y < int(tileQty); y++ {
			freeTiles = append(freeTiles, Position{int32(x), int32(y)} )
		}
	}
}



func setInitialConfig() {
	player = Player{
		body: append([]Position{}, initialBody...),
		dir: Position{x: 0, y: 0},
		
	}
}

func drawSnake(){
	for i := 0; i < len(player.body); i++ {
		body := player.body[i]		
		
		drawSquare(body.x * SquareSize, body.y * SquareSize, SquareSize, SquareSize, rl.DarkGreen)
		
	}
}

func drawApple(){
	drawSquare(apple.x * SquareSize, apple.y * SquareSize, SquareSize , SquareSize, rl.Red)
}

func drawSquare(posX int32, posY int32, sizeX int32, sizeY int32, color rl.Color){
	lineSize := int32(1)
	rl.DrawRectangle(posX + lineSize, posY+ lineSize, sizeX - lineSize * 2, sizeY - lineSize *2, color)

}

