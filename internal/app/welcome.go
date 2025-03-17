package app

import (
	"time"

	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

func welcome(screen *termy.Termy, animate bool) {
	printFunc := typewriter.Write
	delay := 0 * time.Millisecond
	if animate {
		printFunc = typewriter.Type
		delay = 500 * time.Millisecond
	}

	screen.ClearScreen()
	screen.SaveCurPos()
	screen.SetFgHex(palette.Grey)
	screen.Send()

	typewriter.Write("Mec-Nyan's Kana-Master v0.1.0-beta")
	screen.CurToCol(1)
	screen.MoveDown(1)
	typewriter.Write("License information and stuff will be here.")

	screen.SetFgHex(palette.Purple)
	screen.Send()

	screen.MoveTo(8, 4)
	printFunc("  Welcome to Kana-master!")
	time.Sleep(delay)

	screen.MoveTo(8, 6)
	printFunc("Learn ひらがな and カタカナ from the command line.")
	time.Sleep(delay)

	screen.SetFgHex(palette.Blue)
	screen.Italics()
	screen.Send()

	screen.MoveTo(8, 10)
	printFunc("To continue, any key you must press...")

	input.GetChar()
}
