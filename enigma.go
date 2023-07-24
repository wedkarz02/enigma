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

	fonts := gui.InitFonts()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		// rl.DrawText("Enigma", 250, 350, 50, rl.White)
		rl.DrawTextEx(fonts["jbmr"],
			"Enigma",
			rl.Vector2{
				X: 250, Y: 350,
			},
			48, 0,
			rl.White,
		)

		rl.EndDrawing()
	}
}
