package app

import (
	"time"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/rounds"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

func MainLoop(screen *termy.Termy, opts internal.CLIOptions) error {
	defer endMain(screen)

	var quit bool

	if opts.ShowWelcomeScreen {
		quit, err := welcome(screen, opts.Animate)
		if err != nil {
			return err
		}
		if quit {
			return nil
		}
	}

	var userOptions internal.UserOptions

	if opts.ShowOptionsScreen {
		userOptions, quit = setOptions(screen)
		if quit {
			return nil
		}
	}

	for {
		round, quit := rounds.SelectRound(screen, userOptions)
		if quit {
			return nil
		}

		rounds.Fight(screen, round, userOptions)
	}
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
