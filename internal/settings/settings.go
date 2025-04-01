package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

type Option struct {
	Name     string
	Kind     string
	On       bool
	Choice   []string
	Selected int
}

type OptionsMenu []Option

func SetOptions(display *termy.Display) (internal.UserOptions, internal.Action) {
	display.ClearScreen()
	display.HideCur()
	defer display.ShowCur()

	rows, cols, _ := display.Size()
	title := "Options"

	// TODO: Find a better way!
	opts := OptionsMenu{
		{
			Name: "Animations",
			Kind: "bool",
			On:   true,
		},
		{
			Name: "Custom palette",
			Kind: "bool",
			On:   false,
		},
		{
			Name:   "Practise",
			Kind:   "mult",
			Choice: []string{"Hiragana", "Katakana", "Pairs"}},
		{
			Name: "<- Back to main menu",
			Kind: "action",
		},
		{
			Name: "Quit ->",
			Kind: "action",
		},
	}

	var padding int = (cols - len(title)) / 2
	display.MoveTo(padding, 4)
	display.SetFg(4)
	display.Send()
	typewriter.Write(title)

	display.Normal()

	var y int = rows / 3

	current := 0
	exitMsg := "Press q to exit, space to toggle option, enter to accept and continue."
Loop:
	for {
		y := y
		for i, opt := range opts {
			text := fmt.Sprintf("%-20s", opt.Name)
			var toggle string
			if opt.Kind == "bool" {
				toggle = fmt.Sprintf("%20s", onOff(opt.On))
			} else if opt.Kind == "mult" {
				toggle = fmt.Sprintf("%20s", opt.Choice[opt.Selected])
			} else {
				toggle = fmt.Sprintf("%20s", "---")
			}

			if current == i {
				display.SetFg(2)
			} else {
				display.SetFg(5)
			}
			display.Send()

			padding = (cols - (len(text) + len(toggle))) / 2
			display.MoveTo(padding, y)
			typewriter.Write(text)
			// Quit and Back don't toggle.
			if opt.Kind == "bool" && !opt.On {
				display.SetFg(8)
			} else if opt.Kind == "mult" && current == i {
				display.SetFg(3)
			}
			display.Send()
			typewriter.Write(toggle)
			y += 3
		}

		display.MoveTo((cols-len(exitMsg))/2, rows-8)
		display.SetFg(4)
		display.Send()
		os.Stdout.WriteString(exitMsg)

		action, _ := input.GetChar()
		switch action {

		case 'q':
			return internal.UserOptions{}, internal.Quit
		case '\x1b':
			return internal.UserOptions{}, internal.Welcome
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
			if opts[current].Kind == "bool" {
				opts[current].On = !opts[current].On
			} else if opts[current].Kind == "mult" {
				opts[current].Selected++
				if opts[current].Selected == len(opts[current].Choice) {
					opts[current].Selected = 0
				}
			}
		case '\n':
			name := strings.ToLower(opts[current].Name)
			if strings.Contains(name, "quit") {
				return internal.UserOptions{}, internal.Quit
			}
			if strings.Contains(name, "back") {
				return internal.UserOptions{}, internal.Welcome
			}
			break Loop
		}
	}

	// TODO: Better, but still needs improvement!
	return internal.UserOptions{
		Animate:          opts[0].On,
		CustomPalette:    opts[1].On,
		PractisePairs:    opts[2].Choice[opts[2].Selected] == "Pairs",
		PractiseHiragana: opts[2].Choice[opts[2].Selected] == "Hiragana",
		PractiseKatakana: opts[2].Choice[opts[2].Selected] == "Katakana",
	}, internal.Welcome
}

func onOff(value bool) string {
	if value {
		return "( *)  On"
	}
	return "(* ) Off"
}
