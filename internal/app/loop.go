package app

import (
	"time"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/rounds"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/kana-master/pkg/kana"
	"github.com/mec-nyan/termy"
)

type menu struct {
	items map[string]bool
	order []string
}

func MainLoop(screen *termy.Termy, opts internal.CLIOptions) error {
	defer endMain(screen)

	_menu := menu{
		items: map[string]bool{
			"Play":            true,
			"Settings":        false,
			"Round selection": false,
			"Quit":            false,
		},
		order: []string{
			"Play",
			"Settings",
			"Round selection",
			"Quit",
		},
	}

	// This just kept for developping to allow to skip a screen.
	if opts.ShowWelcomeScreen {
		_menu, err := welcome(screen, _menu, opts.Animate)
		if err != nil {
			return err
		}
		if _menu.items["Quit"] {
			return nil
		}
	}

	var quit bool
	var userOptions internal.UserOptions
	round := kana.Rows[0]

	for {
		// TODO: Go straight to the game. But where?
		if _menu.items["Play"] {
			quit = rounds.Fight(screen, round, userOptions)
			if quit {
				return nil
			}
		}
		if _menu.items["Settings"] {
			userOptions, quit = setOptions(screen)
			if quit {
				return nil
			}
		}
		if _menu.items["Round selection"] {
			round, quit = rounds.SelectRound(screen, userOptions)
			if quit {
				return nil
			}
		}
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
