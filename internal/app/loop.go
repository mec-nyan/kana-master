package app

import (
	"time"

	"github.com/mec-nyan/termy"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/help"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/play"
	"github.com/mec-nyan/kana-master/internal/quit"
	"github.com/mec-nyan/kana-master/internal/rounds"
	"github.com/mec-nyan/kana-master/internal/settings"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/kana-master/internal/welcome"

	"github.com/mec-nyan/kana-master/pkg/kana"
)

func MainLoop(display *termy.Display, opts internal.UserOptions) error {
	defer endMain(display, opts)

	round := kana.Rows[0]
	var err error

	// The first Action is always to show the welcome screen.
	action := internal.Welcome
	var previous internal.Action
	var saved internal.Action

	for {
		// TODO: Maybe use channels.
		saved = previous
		previous = action

		switch action {
		case internal.Welcome:
			action, err = welcome.Welcome(display, opts)
		case internal.Play:
			action, err = play.Fight(display, round, opts)
		case internal.Settings:
			opts, action = settings.SetOptions(display)
		case internal.Select:
			round, action = rounds.Selection(display, opts)
		case internal.Continue:
			// TODO: When we reach the end, present an ending screen.
			// Maybe continue to next stage (i.e. from "Hiragana" to
			// "Katakana", from "rows" to "columns", etc).
			round = rounds.NextRow(round)
			action = internal.Play
		case internal.Help:
			action = help.Help(display)
		case internal.Progress:
			return nil
		case internal.Quit:
			action, err = quit.Confirm(display, saved)
		case internal.Exit:
			println("EXIT!")
			return nil
		}

		if err != nil {
			return err
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
