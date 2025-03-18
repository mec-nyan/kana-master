package app

import (
	"fmt"
	"os"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

type Opt struct {
	description string
	value       bool
}

type Quit = bool

func setOptions(screen *termy.Termy, term *termy.TermSettings) (internal.Options, Quit) {
	screen.ClearScreen()
	screen.HideCur()
	defer screen.ShowCur()

	rows, cols, _ := term.Size()
	title := "Options"

	optsOrder := []string{"animation", "palette", "pairs", "hiragana", "katakana"}
	opts := map[string]Opt{
		optsOrder[0]: {"Animations", true},
		optsOrder[1]: {"Custom palette", false},
		optsOrder[2]: {"Practice pairs", true},
		optsOrder[3]: {"Practice hiragana", false},
		optsOrder[4]: {"Practice katakana", false},
	}

	screen.Bold()
	screen.Send()

	var padding int = (cols - len(title)) / 2
	screen.MoveTo(padding, 4)
	screen.SetFgHex(palette.Blue)
	screen.Send()
	typewriter.Write(title)

	screen.Normal()
	screen.SetFgHex(palette.Grey)
	screen.Send()

	var y int = rows / 3

	current := 0
	exitMsg := "Press q to exit, space to toggle option, enter to accept."
Loop:
	for {
		y := y
		for i, key := range optsOrder {
			opt := opts[key]
			text := fmt.Sprintf("%-20s [%3s]", opt.description, onOff(opt.value))
			if current == i {
				screen.SetFg(2)
			} else {
				screen.SetFg(4)
			}
			screen.Send()

			padding = (cols - len(text)) / 2
			screen.MoveTo(padding, y)
			typewriter.Write(text)
			y += 2
		}

		screen.MoveTo((cols-len(exitMsg))/2, rows-8)
		screen.SetFg(8)
		screen.Send()
		os.Stdout.WriteString(exitMsg)

		action, _ := input.GetChar()
		switch action {

		case 'q':
			return internal.Options{}, true
		case '\x1b':
			break Loop
		case 'j', 'n':
			current++
			if current == len(opts) {
				current = 0
			}
		case 'k', 'p':
			current--
			if current < 0 {
				current = len(opts) - 1
			}
		case ' ':
			currenOption := opts[optsOrder[current]]
			currenOption.value = !currenOption.value
			opts[optsOrder[current]] = currenOption
		case '\n':
			break Loop
		}
	}

	return internal.Options{
		AnimationOn:  opts["animation"].value,
		UsePalette:   opts["palette"].value,
		UsePairs:     opts["pairs"].value,
		HiraganaOnly: opts["hiragana"].value,
		KatakanaOnly: opts["katakana"].value,
	}, false
}

func onOff(value bool) string {
	if value {
		return "on"
	}
	return "off"
}
