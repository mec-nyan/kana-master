package app

import (
	"fmt"
	"os"

	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

type Options struct {
	animationOn  bool
	usePalette   bool
	usePairs     bool
	hiraganaOnly bool
	katakanaOnly bool
}

type Opt struct {
	description string
	value       bool
}

type Quit bool

func setOptions(screen *termy.Termy, term *termy.TermSettings) (Options, Quit) {
	screen.ClearScreen()
	screen.HideCur()
	defer screen.ShowCur()

	rows, cols, _ := term.Size()
	title := "Options"

	opts := []Opt{
		{"Animations", true},
		{"Custom palette", false},
		{"Practice pairs", true},
		{"Practice hiragana", false},
		{"Practice katakana", false},
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
		for i, opt := range opts {
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
			return Options{}, true
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
			opts[current].value = !opts[current].value
		case '\n':
			break Loop
		}
	}
	return Options{}, false
}

func onOff(value bool) string {
	if value {
		return "on"
	}
	return "off"
}
