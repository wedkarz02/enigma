package main

import (
	c "enigma/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(c.WINDOW_WIDTH, c.WINDOW_HEIGHT, "Enigma")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Enigma", 250, 350, 50, rl.White)

		rl.EndDrawing()
	}
}
