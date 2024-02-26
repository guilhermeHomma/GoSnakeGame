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

var SquareSize int32
var GameSpeed int
var FrameCount *int 

var player = Player{
	body: []Position{
		{4, 4}, {4,3}, {4,2},
	},
	dir: Position{x: 0, y: 0},
	
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
	if !(*FrameCount % (60 / GameSpeed) == 0 && player.dir != Position{x: 0, y: 0}){
		return
	}

	for i := len(player.body) -1; i >= 0 ; i-- {
		if i == 0 {
			player.body[i].x += player.dir.x
			player.body[i].y += player.dir.y
			continue 
		}
		player.body[i].x = player.body[i - 1].x
		player.body[i].y = player.body[i - 1].y
	}
}

func Update() {
	playerDirection()
	PlayerPosition()
	for i := 0; i < len(player.body); i++ {
		rl.DrawRectangle(player.body[i].x * SquareSize, player.body[i].y * SquareSize, SquareSize, SquareSize, rl.DarkGreen)
	}
}
