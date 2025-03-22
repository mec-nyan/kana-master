package rounds

import (
	"errors"
	"fmt"
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
	menu   internal.ModeMenu
	margin int
}

type RoundMode struct {
	syllabary, group internal.Mode
}

// SelectRound presents a screen with information about completed
// rounds, score, overall, and lets the user select where to go
// from here.
func SelectRound(screen *termy.Termy, opts internal.UserOptions) (
	kana.KanaRow, internal.Action,
) {
	writeFunc := typewriter.Write
	// if opts.AnimationOn {
	// 	writeFunc = typewriter.Type
	// }

	screen.ClearScreen()
	screen.UseDefault()
	screen.Send()

	putCenteredAt(screen, "[( Round Selection )]", 2, writeFunc)

	rows, _, _ := screen.Size()
	xPos, yPos := 4, 4

	index := 1
	for _, row := range kana.Rows {
		if yPos > rows-4 {
			yPos = 4
			xPos += 24
		}
		screen.MoveTo(xPos, yPos)
		writeFunc(fmt.Sprintf("%2d: [ ", index))
		for _, k := range row {
			writeFunc(string(k.Hiragana) + " ")
		}
		writeFunc(" ]")
		yPos += 2
		index++
	}

	putCenteredAt(screen, "Enter a number: ", rows-2, writeFunc)

	i, err := input.GetNumber(3)
	if err != nil || i == -1 || i > len(kana.Rows) {
		return kana.KanaRow{}, internal.Quit
	}

	return kana.Rows[i-1], internal.Play
}

func SelectMode(screen *termy.Termy, _ internal.UserOptions) (RoundMode, internal.Action) {

	syllabary := selectionScreen(screen, MenuScreen{
		title: "Select mode",
		// TODO: How do we include/exclude dakuten/handakuten?
		menu: internal.ModeMenu{
			{
				Name:        "Hiragana",
				Mode:        internal.HiraganaMode,
				Description: "Practise Hiragana",
			},
			{
				Name:        "Katakana",
				Mode:        internal.KatakanaMode,
				Description: "Practise Katakana",
			},
			{
				Name:        "Pairs",
				Mode:        internal.PairsMode,
				Description: "Practise with pairs of Hiragana and Katakana",
			},
		},
		margin: 6,
	})

	if syllabary == internal.NoMode {
		return RoundMode{}, internal.Back
	}

	group := selectionScreen(screen, MenuScreen{
		title: "Select mode",
		// TODO: How do we include/exclude dakuten/handakuten?
		menu: internal.ModeMenu{
			{
				Name:        "Rows",
				Mode:        internal.RowMode,
				Description: "五十音 (gojûon) rows",
			},
			{
				Name:        "Columns",
				Mode:        internal.ColMode,
				Description: "五十音 (gojûon) columns",
			},
			{
				Name:        "Groups",
				Mode:        internal.GroupMode,
				Description: "五十音 (gojûon),　濁点 (dakuten) & 半濁点 (handakuten)",
			},
			{
				Name:        "All",
				Mode:        internal.AllMode,
				Description: "All 五十音 (gojûon)",
			},
		},
		margin: 6,
	})

	if group == internal.NoMode {
		return RoundMode{}, internal.Back
	}

	return RoundMode{syllabary: syllabary, group: group}, internal.Continue
}

func selectionScreen(screen *termy.Termy, menu MenuScreen) internal.Mode {
	screen.ClearScreen()
	screen.UseDefault()
	screen.HideCur()
	screen.Send()
	defer screen.ShowCur()

	rows, cols, _ := screen.Size()

	yPos := menu.margin
	screen.SetFg(6)
	screen.Send()
	screen.MoveTo(1, yPos)
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
				screen.SetFg(2)
			} else {
				screen.SetFg(5)
			}
			screen.Send()

			screen.MoveTo(1, yPos)
			typewriter.WriteCentered("[( "+item.Name+" )]", cols)
			yPos += menuSep
		}

		screen.SetFg(6)
		screen.Send()
		screen.MoveTo(1, rows-menu.margin)
		screen.ClearToEOL()
		typewriter.WriteCentered(menu.menu[selected].Description, cols)

		res, _ := input.GetChar()
		switch res {
		case 'q':
			return internal.NoMode
		case '\x1b':
			return internal.NoMode
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
			return menu.menu[selected].Mode
		}
	}
}

func putCenteredAt(screen *termy.Termy, text string, at int, write func(string)) error {
	rows, cols, err := screen.Size()
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

	screen.MoveTo(xPos, at)
	write(text)

	return nil
}
