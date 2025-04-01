package help

import (
	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

type KeyBinding struct {
	Bindings    []string
	Description string
}

const (
	ShowKeys internal.Action = iota
	ShowHowTo
	ShowAbout
)

func Help(display *termy.Display) internal.Action {
	for {
		action := Select(display)
		switch action {
		case internal.Back:
			return internal.Welcome
		case ShowKeys:
			action = Keybindings(display)
		case ShowHowTo:
			action = GameHelp(display)
		case ShowAbout:
			action = About(display)
		default:
			return internal.Quit
		}
	}
}

func Select(display *termy.Display) internal.Action {
	display.ClearScreen()
	display.HideCur()
	defer display.ShowCur()

	menu := internal.Menu{
		{
			Name:  "Game help",
			Value: ShowHowTo,
		},
		{
			Name:  "Key bindings",
			Value: ShowKeys,
		},
		{
			Name:  "About",
			Value: ShowAbout,
		},
		{
			Name:  "Back",
			Value: internal.Back,
		},
	}
	nameWidth := menu.MaxNameLen()

	rows, cols, _ := display.Size()
	display.SetFg(8)
	display.Send()
	display.MoveTo(1, rows-2)
	typewriter.WriteCentered("Press 'ESC' to go back, 'q' to quit", cols)

	display.MoveTo(1, 4)
	display.SetFg(6)
	display.Send()
	typewriter.WriteCentered("Help", cols)

	yPos := 14
	selected := 0
	for {
		incr := 0
		for i, item := range menu {
			if i == selected {
				display.SetFg(2)
			} else {
				display.SetFg(4)
			}
			display.Send()
			display.MoveTo(1, yPos+incr)
			inner, _ := typewriter.CenterStr(item.Name, nameWidth)
			typewriter.WriteCentered(
				"[( "+inner+" )]", cols,
			)
			incr += 3
		}

		res, _ := input.GetChar()
		switch res {
		case 'q':
			return internal.Quit
		case '\x1b':
			return internal.Welcome
		case 'j', 'n':
			if selected < len(menu)-1 {
				selected++
			}
		case 'k', 'p':
			if selected > 0 {
				selected--
			}
		case '\n':
			return menu[selected].Value
		}
	}
}

func Keybindings(display *termy.Display) internal.Action {
	return wip(display, "Keys")
}

func GameHelp(display *termy.Display) internal.Action {
	return wip(display, "HowTo")
}

func About(display *termy.Display) internal.Action {
	return wip(display, "About")
}

func wip(display *termy.Display, msg string) internal.Action {
	display.ClearScreen()
	display.SetFg(3)
	display.Send()

	rows, cols, _ := display.Size()

	display.MoveTo(1, rows/3)
	typewriter.WriteCentered("[(   WIP: "+msg+"   )]", cols)
	display.MoveTo(1, rows/3+4)
	typewriter.WriteCentered("Press any key...", cols)

	input.GetChar()
	return internal.Back
}

func kb() {
	_ = []KeyBinding{
		{
			[]string{"j", "n"},
			"Down/Next",
		},
		{
			[]string{"k", "p"},
			"Up/Previous",
		},
		{
			[]string{"Enter"},
			"Accept/Continue",
		},
		{
			[]string{"ESC"},
			"Back",
		},
		{
			[]string{"q"},
			"Quit",
		},
	}

}
