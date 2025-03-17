package app

import (
	"time"

	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/rounds"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

func MainLoop(screen *termy.Termy, term *termy.TermSettings) error {
	defer endMain(screen)

	_, quit := setOptions(screen, term)
	if quit {
		return nil
	}

	welcome(screen)

	rounds.Round1Fight(screen)


	return nil
}

func endMain(screen *termy.Termy) {
	screen.ClearScreen()
	screen.SetFgHex(palette.Green)
	screen.Send()
	screen.MoveTo(4, 2)
	typewriter.Write("Bye!")
	time.Sleep(time.Second * 1)
	screen.ClearScreen()
}

func ClearFromSavedPos(screen *termy.Termy) {
	screen.RestoreCurPos()
	screen.ClearToEOS()
}
