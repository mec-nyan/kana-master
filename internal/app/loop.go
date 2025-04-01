package app

import (
	"time"

	"github.com/mec-nyan/termy"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/help"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/play"
	"github.com/mec-nyan/kana-master/internal/rounds"
	"github.com/mec-nyan/kana-master/internal/settings"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/kana-master/internal/welcome"

	"github.com/mec-nyan/kana-master/pkg/kana"
)

func MainLoop(display *termy.Display, opts internal.UserOptions) error {
	defer endMain(display, opts)

	var round kana.KanaRow
	var err error

	// The first Action is always to show the welcome screen.
	action := internal.Welcome

	for {
		if action == internal.Welcome {
			action, err = welcome.Welcome(display, opts)
			if err != nil {
				return err
			}
		} else if action == internal.Play {
			// The user hasn't selected a round yet, use the first one.
			if round == nil {
				round = kana.Rows[0]
			}
			action, err = play.Fight(display, round, opts)
		} else if action == internal.Settings {
			opts, action = settings.SetOptions(display)
		} else if action == internal.Select {
			round, action = rounds.Selection(display, opts)
		} else if action == internal.Continue {
			// TODO: When we reach the end, present an ending screen.
			// Maybe continue to next stage (i.e. from "Hiragana" to
			// "Katakana", from "rows" to "columns", etc).
			round = rounds.NextRow(round)
			action = internal.Play
		} else if action == internal.Help {
			action = help.Help(display)
		} else if action == internal.Progress {
			return nil
		} else if action == internal.Quit {
			return nil
		}
	}
}

func endMain(display *termy.Display, opts internal.UserOptions) {
	// TODO: Add some cool ASCII art to the end screen 💖
	writeFunc := typewriter.Write
	if opts.Animate {
		writeFunc = typewriter.Type
	}

	display.ClearScreen()
	display.SetFgHex(palette.Green)
	display.Send()

	rows, cols, err := display.Size()
	if err != nil {
		// TODO: How to handle this error?
		panic(err)
	}

	bye := "Bye!"
	y := rows / 3
	x := (cols - len(bye)) / 2
	display.MoveTo(x, y)

	writeFunc(bye)

	time.Sleep(time.Second * 1)
	display.ClearScreen()
}

func ClearFromSavedPos(display *termy.Display) {
	display.RestoreCurPos()
	display.ClearToEOS()
}
