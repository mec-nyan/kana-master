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
	menu   internal.RoundModeMenu
	margin int
}

// SelectRound presents a screen with information about completed
// rounds, score, overall, and lets the user select where to go
// from here.
func SelectRound(screen *termy.Termy, mode RoundMode, opts internal.UserOptions) (
	kana.KanaRow, internal.Action,
) {
	writeFunc := typewriter.Write
	// if opts.AnimationOn {
	// 	writeFunc = typewriter.Type
	// }

	screen.ClearScreen()
	screen.UseDefault()
	screen.SetFg(1)
	screen.Send()

	putCenteredAt(screen, "Round Selection", 2, writeFunc)

	rows, cols, _ := screen.Size()

	rounds := getRounds(mode.Group)

	yPos := 6
	selected := 0
	for {
		yPos := yPos
		for i, round := range rounds {
			if i == selected {
				screen.SetFg(2)
			} else {
				screen.SetFg(5)
			}
			screen.Send()
			screen.MoveTo(1, yPos)
			yPos += 2

			if mode.Syllabary == internal.HiraganaMode {
				if round[0].Romaji == "ya" || round[0].Romaji == "wa" {
					typewriter.WriteCentered(
						fmt.Sprintf("%c    %c    %c",
							round[0].Hiragana,
							round[2].Hiragana,
							round[4].Hiragana,
						),
						cols,
					)
				} else {
					typewriter.WriteCentered(
						fmt.Sprintf("%c %c %c %c %c",
							round[0].Hiragana,
							round[1].Hiragana,
							round[2].Hiragana,
							round[3].Hiragana,
							round[4].Hiragana,
						),
						cols,
					)
				}
			} else if mode.Syllabary == internal.KatakanaMode {
				if round[0].Romaji == "ya" || round[0].Romaji == "wa" {
					typewriter.WriteCentered(
						fmt.Sprintf("%c    %c    %c",
							round[0].Katakana,
							round[2].Katakana,
							round[4].Katakana,
						),
						cols,
					)
				} else {
					typewriter.WriteCentered(
						fmt.Sprintf("%c %c %c %c %c",
							round[0].Katakana,
							round[1].Katakana,
							round[2].Katakana,
							round[3].Katakana,
							round[4].Katakana,
						),
						cols,
					)
				}
			} else if mode.Syllabary == internal.PairsMode {
				typewriter.WriteCentered(
					fmt.Sprintf("(%c, %c), (%c, %c), (%c, %c), (%c, %c), (%c, %c)",
						round[0].Hiragana, round[0].Katakana,
						round[1].Hiragana, round[1].Katakana,
						round[2].Hiragana, round[2].Katakana,
						round[3].Hiragana, round[3].Katakana,
						round[4].Hiragana, round[4].Katakana,
					),
					cols,
				)
			}
		}

		screen.MoveTo(1, rows-2)
		screen.SetFg(5)
		screen.Send()
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

func SelectSyllabary(screen *termy.Termy, _ internal.UserOptions) (internal.RoundMode, internal.Action) {

	syllabary := selectionScreen(screen, MenuScreen{
		title: "Select mode",
		// TODO: How do we include/exclude dakuten/handakuten?
		menu: internal.RoundModeMenu{
			{
				Name:        "Hiragana",
				RoundMode:   internal.HiraganaMode,
				Description: "Practise Hiragana",
			},
			{
				Name:        "Katakana",
				RoundMode:   internal.KatakanaMode,
				Description: "Practise Katakana",
			},
			{
				Name:        "Pairs",
				RoundMode:   internal.PairsMode,
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

func SelectGroup(screen *termy.Termy, _ internal.UserOptions) (internal.RoundMode, internal.Action) {

	group := selectionScreen(screen, MenuScreen{
		title: "Select mode",
		// TODO: How do we include/exclude dakuten/handakuten?
		menu: internal.RoundModeMenu{
			{
				Name:        "Rows",
				RoundMode:   internal.RowMode,
				Description: "五十音 (gojûon) rows",
			},
			{
				Name:        "Columns",
				RoundMode:   internal.ColMode,
				Description: "五十音 (gojûon) columns",
			},
			{
				Name:        "Groups",
				RoundMode:   internal.GroupMode,
				Description: "五十音 (gojûon),　濁点 (dakuten) & 半濁点 (handakuten)",
			},
			{
				Name:        "All",
				RoundMode:   internal.AllMode,
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

func selectionScreen(screen *termy.Termy, menu MenuScreen) internal.RoundMode {
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
			return menu.menu[selected].RoundMode
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
