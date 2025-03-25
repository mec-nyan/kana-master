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

func Help(screen *termy.Termy) internal.Action {
	for {
		action := Select(screen)
		switch action {
		case internal.Back:
			return internal.Welcome
		case ShowKeys:
			action = Keybindings(screen)
		case ShowHowTo:
			action = GameHelp(screen)
		case ShowAbout:
			action = About(screen)
		default:
			return internal.Quit
		}
	}
}

func Select(screen *termy.Termy) internal.Action {
	screen.ClearScreen()
	screen.HideCur()
	defer screen.ShowCur()

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

	rows, cols, _ := screen.Size()
	screen.SetFg(8)
	screen.Send()
	screen.MoveTo(1, rows-2)
	typewriter.WriteCentered("Press 'ESC' to go back, 'q' to quit", cols)

	screen.MoveTo(1, 4)
	screen.SetFg(6)
	screen.Send()
	typewriter.WriteCentered("Help", cols)

	yPos := 14
	selected := 0
	for {
		incr := 0
		for i, item := range menu {
			if i == selected {
				screen.SetFg(2)
			} else {
				screen.SetFg(4)
			}
			screen.Send()
			screen.MoveTo(1, yPos+incr)
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

func Keybindings(screen *termy.Termy) internal.Action {
	return wip(screen, "Keys")
}

func GameHelp(screen *termy.Termy) internal.Action {
	return wip(screen, "HowTo")
}

func About(screen *termy.Termy) internal.Action {
	return wip(screen, "About")
}

func wip(screen *termy.Termy, msg string) internal.Action {
	screen.ClearScreen()
	screen.SetFg(3)
	screen.Send()

	rows, cols, _ := screen.Size()

	screen.MoveTo(1, rows/3)
	typewriter.WriteCentered("[(   WIP: "+msg+"   )]", cols)
	screen.MoveTo(1, rows/3+4)
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
