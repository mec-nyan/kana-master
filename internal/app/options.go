package app

import (
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

func options(screen *termy.Termy, term *termy.TermSettings) {
	screen.ClearScreen()

	rows, cols, _ := term.Size()
	title := "Options"

	opts := []string{
		"[-]   Use app's palette    ",
		"[ ]   Use terminal colours ",
		"[ ]   This is actually just",
		"[ ]   a placeholder. Here  ",
		"[ ]   we'll be real options",
		"[ ]   soon!                ",
		"[ ]   Just press any key   ",
		"[ ]          to continue...",
	}

	screen.Bold()
	screen.Send()

	var padding int = (cols - len(title)) / 2
	screen.MoveTo(padding, 4)
	screen.SetFgHex(palette.Blue)
	screen.Send()
	typewriter.Write(title)

	screen.Normal()
	screen.SetFgHex(palette.Grey)
	screen.Send()

	var y int = rows / 3

	for _, opt := range opts {
		padding = (cols - len(opt)) / 2
		screen.MoveTo(padding, y)
		typewriter.Write(opt)
		y += 2
	}

	input.GetChar()
}
