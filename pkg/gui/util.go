package gui

import (
	"log"

	c "enigma/pkg/consts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fonts map[string]rl.Font

type Text struct {
	Content  string
	Font     rl.Font
	FontSize int
	Spacing  float32
	Color    rl.Color
	Active   bool
}

func (text *Text) Size() rl.Vector2 {
	return rl.MeasureTextEx(
		text.Font,
		text.Content,
		float32(text.FontSize),
		text.Spacing,
	)
}

func (list Fonts) LoadFont(name string, alias string, defRes int32) {
	if _, inFonts := list[alias]; inFonts {
		log.Fatalf("[ERROR]: %v already loaded\n", name)
	}

	list[alias] = rl.LoadFontEx(c.FONTS_PATH+name, defRes, nil)
}

func InitFonts() Fonts {
	fonts := make(Fonts)

	fonts.LoadFont("JetBrainsMono-Regular.ttf", "jbmr", c.KEY_SIZE)

	return fonts
}
