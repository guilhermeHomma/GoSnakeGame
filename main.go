package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"SnakeGame/game"
)

const (
	screenW = 720
	screenH = 720
)


var frameCount int = 0
var squareSize int = 45
var gameSpeed int = 7

func main() {

	rl.InitWindow(screenW, screenH, "test")
	rl.SetTargetFPS(60)
	
	game.SquareSize = int32(squareSize)
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

