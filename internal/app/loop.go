package app

import (
	"time"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/rounds"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

func MainLoop(screen *termy.Termy, term *termy.TermSettings, opts InitOptions) error {
	defer endMain(screen)

	if opts.Welcome {
		welcome(screen, term, opts.Animate)
	}

	var quit bool
	var userOptions internal.Options

	if opts.Options {
		userOptions, quit = setOptions(screen, term)
		if quit {
			return nil
		}
	}

	rounds.Round1Fight(screen, userOptions)

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
