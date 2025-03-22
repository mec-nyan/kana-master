package app

import (
	"time"

	"github.com/mec-nyan/termy"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/options"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/play"
	"github.com/mec-nyan/kana-master/internal/rounds"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/kana-master/internal/welcome"

	"github.com/mec-nyan/kana-master/pkg/kana"
)

func MainLoop(screen *termy.Termy, opts internal.CLIOptions) error {
	defer endMain(screen, opts)

	var round kana.KanaRow
	var quit bool
	var err error
	var userOptions internal.UserOptions

	// The first Action is always to show the welcome screen.
	action := internal.Welcome
	// mode := internal.RowMode

	for {

		if action == internal.Welcome {
			action, err = welcome.Welcome(screen, opts.Animate)
			if err != nil {
				return err
			}
		} else if action == internal.Play {
			// The user hasn't selected a round yet, use the first one.
			if round == nil {
				round = kana.Rows[0]
			}
			action, err = play.Fight(screen, round, userOptions)
			if quit {
				return nil
			}
		} else if action == internal.Settings {
			userOptions, action = options.SetOptions(screen)
			if action == internal.Back {
				action = internal.Welcome
			}
		} else if action == internal.Select {
			_, action = rounds.SelectMode(screen, userOptions)
		} else if action == internal.Continue {
			// TODO: When we reach the end, present an ending screen.
			// Maybe continue to next stage (i.e. from "Hiragana" to
			// "Katakana", from "rows" to "columns", etc).
			round = rounds.NextRow(round)
			action = internal.Play
		} else if action == internal.Quit {
			return nil
		}
	}
}

func endMain(screen *termy.Termy, opts internal.CLIOptions) {
	// TODO: Add some cool ASCII art to the end screen 💖
	writeFunc := typewriter.Write
	if opts.Animate {
		writeFunc = typewriter.Type
	}

	screen.ClearScreen()
	screen.SetFgHex(palette.Green)
	screen.Send()

	rows, cols, err := screen.Size()
	if err != nil {
		// TODO: How to handle this error?
		panic(err)
	}

	bye := "Bye!"
	y := rows / 3
	x := (cols - len(bye)) / 2
	screen.MoveTo(x, y)

	writeFunc(bye)

	time.Sleep(time.Second * 1)
	screen.ClearScreen()
}

func ClearFromSavedPos(screen *termy.Termy) {
	screen.RestoreCurPos()
	screen.ClearToEOS()
}
