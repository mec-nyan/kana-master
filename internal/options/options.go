package options

import (
	"fmt"
	"os"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

type options struct {
	items map[string]bool
	order []string
}

func SetOptions(screen *termy.Termy) (internal.UserOptions, internal.Action) {
	screen.ClearScreen()
	screen.HideCur()
	defer screen.ShowCur()

	rows, cols, _ := screen.Size()
	title := "Options"

	// TODO: Find a better way!
	opts := options{
		items: map[string]bool{
			"Animations":        true,
			"Custom palette":    false,
			"Practise pairs":    true,
			"Practise hiragana": false,
			"Practise katakana": false,
			"Back to main menu": false,
			"Quit":              false,
		},
		order: []string{
			"Animations",
			"Custom palette",
			"Practise pairs",
			"Practise hiragana",
			"Practise katakana",
			"Back to main menu",
			"Quit",
		},
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
		for i, opt := range opts.order {
			var text string
			if opt != "Back to main menu" && opt != "Quit" {
				text = fmt.Sprintf("%-20s [%3s]", opt, onOff(opts.items[opt]))
			} else {
				text = fmt.Sprintf("%-26s", opt)
			}
			if current == i {
				screen.SetFg(2)
			} else {
				screen.SetFg(4)
			}
			screen.Send()

			padding = (cols - len(text)) / 2
			screen.MoveTo(padding, y)
			typewriter.Write(text)
			y += 3
		}

		screen.MoveTo((cols-len(exitMsg))/2, rows-8)
		screen.SetFg(8)
		screen.Send()
		os.Stdout.WriteString(exitMsg)

		action, _ := input.GetChar()
		switch action {

		case 'q':
			return internal.UserOptions{}, internal.Quit
		case '\x1b':
			return internal.UserOptions{}, internal.Back
		case 'j', 'n':
			current++
			if current == len(opts.items) {
				current = 0
			}
		case 'k', 'p':
			current--
			if current < 0 {
				current = len(opts.items) - 1
			}
		case ' ':
			opts.items[opts.order[current]] = !opts.items[opts.order[current]]
		case '\n':
			if opts.order[current] == "Quit" {
				return internal.UserOptions{}, internal.Quit
			}
			if opts.order[current] == "Back to main menu" {
				return internal.UserOptions{}, internal.Back
			}
			break Loop
		}
	}

	// TODO: I don't like this!
	return internal.UserOptions{
		AnimationOn:      opts.items["Animations"],
		UsePalette:       opts.items["Custom palette"],
		PractisePairs:    opts.items["PractisePairs"],
		PractiseHiragana: opts.items["Practise hiragana"],
		PractiseKatakana: opts.items["Practise katakana"],
		BackToMain:       opts.items["Back to main menu"],
	}, internal.Welcome
}

func onOff(value bool) string {
	if value {
		return "on"
	}
	return "off"
}
