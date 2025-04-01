package welcome

import (
	"strings"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

var mainMenu = internal.Menu{
	{
		Name:  "Play",
		Value: internal.Play,
		// Description: "Play",
	},
	{
		Name:  "Settings",
		Value: internal.Settings,
		// Description: "Settings",
	},
	{
		Name:  "Round  Selection",
		Value: internal.Select,
		// Description: "Round  Selection",
	},
	{
		Name:  "Help",
		Value: internal.Help,
		// Description: "Help",
	},
	{
		Name:  "Quit",
		Value: internal.Quit,
		// Description: "Quit",
	},
}

// Welcome presents the Welcome and selection screen.
func Welcome(display *termy.Display, opts internal.UserOptions) (internal.Action, error) {
	display.HideCur()
	defer display.ShowCur()

	// Every screen should take care of cleaning and setting up the display.
	// We shouldn't take for granted that the previous screen cleaned everything up
	// successfully.
	display.ClearScreen()
	// TODO: Handle colours (default terminal theme, custom palette, etc)
	display.SetFgHex(palette.Grey)
	display.Send()

	putInfo(display)

	// TODO: Check for minimun height and width.
	_, cols, _ := display.Size()

	yPos := 4
	yPos = paintBanner(display, yPos, cols)
	yPos = putTitle(display, yPos, cols)

	yPos += 8

	selected := 0
	var action internal.Action
	for {
		putMenu(display, mainMenu, yPos, cols, selected)
		action, selected, _ = handleInput(selected, mainMenu)
		if action != internal.Continue {
			return action, nil
		}
	}
}

func putInfo(display *termy.Display) {
	// TODO: Select dinamically the correct version.
	typewriter.Write("Mec-Nyan's Kana-Master v0.1.0-beta.")
	display.MoveTo(1, 2)
	typewriter.Write("Kana-Master is released under GPL-3.0 license.")
}

func paintBanner(display *termy.Display, yPos, cols int) int {
	headerLines := strings.Split(banner, "\n")

	display.SetFgHex(palette.Blue)
	display.Send()

	for _, line := range headerLines {
		display.MoveTo(1, yPos)
		typewriter.WriteCentered(line, cols)
		yPos++
	}
	return yPos
}

func putTitle(display *termy.Display, yPos, cols int) int {
	// For now I've remove the typing animations.
	// I'm thinking the best way to add some nice animations later on.
	display.SetFgHex(palette.Purple)
	display.Send()

	yPos += 2
	display.MoveTo(1, yPos)
	typewriter.WriteCentered("    Welcome to Kana-master!    ", cols)

	yPos += 2
	display.MoveTo(1, yPos)
	typewriter.WriteCentered("Learn hiragana and katakana from the command line", cols)

	return yPos
}

func putMenu(display *termy.Display, menu internal.Menu, yPos, cols, sel int) {

	nameWidth := menu.MaxNameLen()
	incr := 0
	// Paint the selection menu.
	for i, item := range menu {
		if i == sel {
			display.SetFg(2)
		} else {
			display.SetFg(4)
		}
		display.Send()
		display.MoveTo(1, yPos+incr)
		inner, _ := typewriter.CenterStr(item.Name, nameWidth)
		typewriter.WriteCentered("[( "+inner+" )]", cols)
		incr += 3
	}
}

func handleInput(selected int, menu internal.Menu) (internal.Action, int, error) {
	esc := false
	csi := false
	for {
		key, err := input.GetChar()
		if err != nil {
			return internal.NoOp, selected, err
		}
		switch key {
		case 'j', 'n':
			selected++
			if selected == len(menu) {
				selected = 0
			}
		case 'k', 'p':
			selected--
			if selected < 0 {
				selected = len(menu) - 1
			}
		case 'q':
			return internal.Quit, selected, nil
		case '\n':
			return menu[selected].Value, selected, nil
		// This is maybe a too complicated way to handle the arrow keys...
		case '\x1b':
			esc = true
			continue
		case '[':
			if esc {
				csi = true
				continue
			}
		case 'B':
			if csi {
				selected++
				if selected == len(menu) {
					selected = 0
				}
				esc = false
				csi = false
			}
		case 'A':
			if csi {
				selected--
				if selected < 0 {
					selected = len(menu) - 1
				}
				esc = false
				csi = false
			}
		}
		return internal.Continue, selected, nil
	}
}
