package rounds

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/kana-master/pkg/kana"
	"github.com/mec-nyan/termy"
)

// Can we do this generic for Actions and Modes?
type MenuScreen struct {
	title  string
	menu   internal.Menu
	margin int
}

func Selection(display *termy.Display, opts internal.UserOptions) (
	kana.KanaRow, internal.Action,
) {
	// First we select what we want to practise.
	action := internal.SelectSyllabary
	var mode RoundMode
	var round kana.KanaRow
	for {
		if action == internal.SelectSyllabary {
			mode.Syllabary, action = SelectSyllabary(display, opts)
			if action == internal.Back {
				return kana.KanaRow{}, internal.Welcome
			}
		} else if action == internal.SelectGroup {
			mode.Group, action = SelectGroup(display, opts)
			if action == internal.Back {
				action = internal.SelectSyllabary
			}
		} else if action == internal.SelectRound {
			round, action = SelectRound(display, mode, opts)
			if action == internal.Back {
				action = internal.SelectGroup
				continue
			}
			return round, action
		}
	}
}

// SelectRound presents a screen with information about completed
// rounds, score, overall, and lets the user select where to go
// from here.
func SelectRound(display *termy.Display, mode RoundMode, opts internal.UserOptions) (
	kana.KanaRow, internal.Action,
) {
	writeFunc := typewriter.Write
	// if opts.AnimationOn {
	// 	writeFunc = typewriter.Type
	// }

	display.ClearScreen()
	display.UseDefault()
	display.SetFg(4)
	display.Send()

	putCenteredAt(display, "Round Selection", 4, writeFunc)

	rows, cols, _ := display.Size()

	rounds := getRounds(mode.Group)

	yPos := (rows - len(rounds)*2) / 2
	selected := 0
	// Event loop.
	for {
		yPos := yPos
		for i, round := range rounds {

			name, group := getRoundLine(round, mode)
			pointer := "  "
			// Double "name" so we center the "group"
			size := kana.CountCols(name + " " + group + pointer)
			padding := (cols - size) / 2
			display.MoveTo(padding, yPos)

			if i == selected {
				display.SetFg(3)
			} else {
				display.SetFg(242)
			}
			display.Send()
			typewriter.Write(name)
			if i == selected {
				display.SetFg(2)
			} else {
				display.SetFg(5)
			}
			display.Send()
			typewriter.Write(group)
			if i == selected {
				display.SetFg(3).Send()
				typewriter.Write(pointer)
			} else {
				typewriter.Write("   ")
			}

			yPos += 2
		}

		// TODO: Check for terminal size!
		display.MoveTo(1, rows-4)
		display.SetFg(4)
		display.Send()
		typewriter.WriteCentered("Press Escape to go back", cols)

		res, _ := input.GetChar()
		switch res {
		case '\x1b':
			return kana.KanaRow{}, internal.Back
		case 'q':
			return kana.KanaRow{}, internal.Quit
		case 'j', 'n':
			selected++
			if selected == len(rounds) {
				selected = 0
			}
		case 'k', 'p':
			selected--
			if selected < 0 {
				selected = len(rounds) - 1
			}
		case '\n':
			return rounds[selected], internal.Play
		}
	}
}

func SelectSyllabary(display *termy.Display, _ internal.UserOptions) (internal.RoundMode, internal.Action) {

	syllabary := selectionScreen(display, MenuScreen{
		title: "Select mode",
		// TODO: How do we include/exclude dakuten/handakuten?
		menu: internal.Menu{
			{
				Name:        "Hiragana",
				Value:       internal.HiraganaMode,
				Description: "Practise Hiragana",
			},
			{
				Name:        "Katakana",
				Value:       internal.KatakanaMode,
				Description: "Practise Katakana",
			},
			{
				Name:        "Pairs",
				Value:       internal.PairsMode,
				Description: "Practise with pairs of Hiragana and Katakana",
			},
		},
		margin: 6,
	})

	if syllabary == 0 {
		return 0, internal.Back
	}

	return syllabary, internal.SelectGroup
}

func SelectGroup(display *termy.Display, _ internal.UserOptions) (internal.RoundMode, internal.Action) {

	group := selectionScreen(display, MenuScreen{
		title: "Select mode",
		// TODO: How do we include/exclude dakuten/handakuten?
		menu: internal.Menu{
			{
				Name:        "Rows",
				Value:       internal.RowMode,
				Description: "五十音 (gojûon) rows",
			},
			{
				Name:        "Columns",
				Value:       internal.ColMode,
				Description: "五十音 (gojûon) columns",
			},
			{
				Name:        "Groups",
				Value:       internal.GroupMode,
				Description: "五十音 (gojûon),　濁点 (dakuten) & 半濁点 (handakuten)",
			},
			{
				Name:        "All",
				Value:       internal.AllMode,
				Description: "All 五十音 (gojûon)",
			},
		},
		margin: 6,
	})

	if group == 0 {
		return 0, internal.Back
	}

	return group, internal.SelectRound
}

func selectionScreen(display *termy.Display, menu MenuScreen) internal.RoundMode {
	display.ClearScreen()
	display.UseDefault()
	display.HideCur()
	display.Send()
	defer display.ShowCur()

	rows, cols, _ := display.Size()
	nameWidth := menu.menu.MaxNameLen()

	yPos := menu.margin
	display.SetFg(4)
	display.Send()
	display.MoveTo(1, yPos)
	typewriter.WriteCentered(menu.title, cols)

	menuSep := 3
	menuHeight := menuSep*(len(menu.menu)-1) + 1

	selected := 0
	for {
		yPos := (rows - menuHeight) / 2
		// Draw the menu.
		for i, item := range menu.menu {
			// Highlight current selected option
			if i == selected {
				display.SetFg(2)
			} else {
				display.SetFg(5)
			}
			display.Send()

			display.MoveTo(1, yPos)
			name, _ := typewriter.CenterStr(item.Name, nameWidth)
			typewriter.WriteCentered("[( "+name+" )]", cols)
			yPos += menuSep
		}

		display.SetFg(4)
		display.Send()
		display.MoveTo(1, rows-menu.margin)
		display.ClearToEOL()
		typewriter.WriteCentered(menu.menu[selected].Description, cols)

		res, _ := input.GetChar()
		switch res {
		case 'q':
			return 0
		case '\x1b':
			return 0
		case 'j', 'n':
			selected++
			if selected == len(menu.menu) {
				selected = 0
			}
		case 'k', 'p':
			selected--
			if selected < 0 {
				selected = len(menu.menu) - 1
			}
		case '\n':
			return menu.menu[selected].Value
		}
	}
}

func getRoundLine(round kana.KanaRow, mode RoundMode) (name, group string) {

	if mode.Group == internal.RowMode {
		name = "[" + strings.ToUpper(string(round[0].Romaji))[:1] + "]"
	} else if mode.Group == internal.ColMode {
		name = "[" + strings.ToUpper(string(round[0].Romaji)) + "]"
	}

	if mode.Syllabary == internal.HiraganaMode {
		for _, n := range round {
			c := n.Hiragana
			if c == "" {
				group += "   "
			} else {
				group += fmt.Sprintf(" %s", c)
			}
		}
	} else if mode.Syllabary == internal.KatakanaMode {
		for _, n := range round {
			c := n.Katakana
			if c == "" {
				group += "   "
			} else {
				group += fmt.Sprintf(" %s", c)
			}
		}
	} else if mode.Syllabary == internal.PairsMode {
		for _, n := range round {
			h := n.Hiragana
			k := n.Hiragana
			if h == "" {
				group += "      "
			} else {
				group += fmt.Sprintf(" %s %s,", h, k)
			}
		}
	}
	return
}

func putCenteredAt(display *termy.Display, text string, at int, write func(string)) error {
	rows, cols, err := display.Size()
	if err != nil {
		return err
	}

	if at > rows {
		return errors.New("Off screen!")
	}

	// TODO: We need to count COLUMNS and not characters,
	// some characters may occupy more than one column.
	size := utf8.RuneCountInString(text)
	if size > cols {
		return errors.New("Too long!")
	}

	xPos := (cols - size) / 2

	display.MoveTo(xPos, at)
	write(text)

	return nil
}
