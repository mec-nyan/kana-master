package app

import (
	"time"

	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

func welcome(screen *termy.Termy) {
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
	typewriter.Type("  Welcome to Kana-master!")
	time.Sleep(500 * time.Millisecond)

	screen.MoveTo(8, 6)
	typewriter.Type("Learn ひらがな and カタカナ from the command line.")
	time.Sleep(500 * time.Millisecond)

	screen.SetFgHex(palette.Blue)
	screen.Italics()
	screen.Send()

	screen.MoveTo(8, 10)
	typewriter.Type("To continue, any key you press must...")

	input.GetChar()
}
