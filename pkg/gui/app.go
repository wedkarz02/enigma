package gui

import (
	c "enigma/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	Fonts    Fonts
	KeySet   Text
	Keyboard []Text
}

func InitApp() *App {
	app := new(App)

	app.Fonts = InitFonts()

	letterSpacing := 48
	app.KeySet = Text{
		Content:  c.KEY_SET,
		Font:     app.Fonts["jbmr"],
		FontSize: c.KEY_SIZE,
		Spacing:  float32(letterSpacing),
		Color:    rl.White,
		Active:   false,
	}

	for _, chr := range c.KEY_SET {
		letter := Text{
			Content:  string(chr),
			Font:     app.Fonts["jbmr"],
			FontSize: c.KEY_SIZE,
			Spacing:  float32(letterSpacing),
			Color:    rl.White,
			Active:   false,
		}

		app.Keyboard = append(app.Keyboard, letter)
	}

	return app
}

func (app *App) DrawKeys() {

}
