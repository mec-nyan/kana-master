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

func MainLoop(screen *termy.Termy, opts internal.CLIOptions) error {
	defer endMain(screen)

	var quit bool
	var err error
	var userOptions internal.UserOptions

	var action internal.Action = "welcome"
	round := kana.Rows[0]

	for {

		if action == "welcome" {
			action, err = welcome(screen, opts.Animate)
			if err != nil {
				return err
			}
		} else if action == "play" {
			action, err = rounds.Fight(screen, round, userOptions)
			if quit {
				return nil
			}
		} else if action == "settings" {
			userOptions, action = setOptions(screen)
			if action == "back" {
				action = "welcome"
			}
		} else if action == "select" {
			round, action = rounds.SelectRound(screen, userOptions)
		} else if action == "continue" {
			// TODO: When we reach the end, present an ending screen.
			// Maybe continue to next stage (i.e. from "Hiragana" to
			// "Katakana", from "rows" to "columns", etc).
			round = rounds.NextRow(round)
			action = "play"

		} else if action == "quit" {
			return nil
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
