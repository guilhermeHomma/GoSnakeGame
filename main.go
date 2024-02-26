package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"SnakeGame/game"
	"time"
	"math/rand"
)

const (
	screenW = 720
	screenH = 720
)


var frameCount int = 0
var squareSize int = 45
var gameSpeed int = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	rl.InitWindow(screenW, screenH, "test")
	rl.SetTargetFPS(60)
	
	game.SquareSize = int32(squareSize)
	game.GameSize = screenW
	game.GameSpeed = gameSpeed
	game.FrameCount = &frameCount

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		frameCount++
		game.Update()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

