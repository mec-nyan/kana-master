package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
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

func SetOptions(screen *termy.Termy) (internal.UserOptions, internal.Action) {
	screen.ClearScreen()
	screen.HideCur()
	defer screen.ShowCur()

	rows, cols, _ := screen.Size()
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
				screen.SetFg(2)
			} else {
				screen.SetFg(4)
			}
			screen.Send()

			padding = (cols - (len(text) + len(toggle))) / 2
			screen.MoveTo(padding, y)
			typewriter.Write(text)
			// Quit and Back don't toggle.
			if opt.Kind == "bool" && !opt.On {
				screen.SetFg(8)
			} else if opt.Kind == "mult" && current == i {
				screen.SetFg(3)
			}
			screen.Send()
			typewriter.Write(toggle)
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
