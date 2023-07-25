package main

import (
	c "enigma/pkg/consts"
	"enigma/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(c.WINDOW_WIDTH, c.WINDOW_HEIGHT, "Enigma")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	app := gui.InitApp()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		app.DrawKeys()

		rl.EndDrawing()
	}
}
