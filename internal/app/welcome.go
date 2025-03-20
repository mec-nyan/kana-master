package app

import (
	"strings"
	"time"

	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

// TODO: Make a cuter header 🐈
const header = `
█   █                      █▀▀█▀▀█                              
█   █  █████ █████ █████   █  █  █ █████ █████ █████ ████ █████ 
██████ █   █ █   █ █   █   ██ █  █ █   █ █   ▀   █   █    █   █ 
██   █ █████ ██  █ █████   ██ █  █ █████ █████   ██  ████ ██████
██   █ ██  █ ██  █ ██  █   ██ █  █ ██  █    ██   ██  ██   ██   █
██   █ ██  █ ██  █ ██  █   ██ █  █ ██  █ █████   ██  ████ ██   █
`

// welcome presents the welcome and selection screen.
func welcome(screen *termy.Termy, _menu menu, animate bool) (menu, error) {
	screen.HideCur()
	defer screen.ShowCur()

	// printFunc := typewriter.Write
	printCenteredFunc := typewriter.WriteCenterd
	delay := 0 * time.Millisecond
	if animate {
		// printFunc = typewriter.Type
		printCenteredFunc = typewriter.TypeCentered
		delay = 500 * time.Millisecond
	}

	// TODO: Check for minimun height and width.
	_, cols, _ := screen.Size()

	// Every screen should take care of cleaning and setting up the display.
	// We shouldn't take for granted that the previous screen cleaned everything up
	// successfully.
	screen.ClearScreen()
	screen.SetFgHex(palette.Grey)
	screen.Send()

	// TODO: Select dinamically the correct version.
	typewriter.Write("Mec-Nyan's Kana-Master v0.1.0-beta.")
	screen.MoveTo(1, 2)
	typewriter.Write("Kana-Master is released under GPL-3.0 license.")

	headerLines := strings.Split(header, "\n")

	// Center the banner.
	headerY := 4
	headerX := 1

	screen.SetFgHex(palette.Blue)
	screen.Send()

	for _, line := range headerLines {
		screen.MoveTo(headerX, headerY)
		typewriter.WriteCenterd(line, cols)
		headerY++
	}

	screen.SetFgHex(palette.Purple)
	screen.Send()

	y := headerY + 2
	screen.MoveTo(1, y)
	// For some reason, these icons are printed fine (right number of cols is detected)
	// but hiragana and katakana aren't.
	printCenteredFunc("    Welcome to Kana-master!    ", cols)
	time.Sleep(delay)

	y += 2
	screen.MoveTo(1, y)
	// TODO: How to count ひらがな and カタカナ columns???
	// printCenteredFunc("Learn ひらがな and カタカナ from the command line.", cols)
	printCenteredFunc("Learn hiragana and katakana from the command line", cols)
	time.Sleep(delay)

	y += 8
	selected := 0
	for {
		incr := 0
		// Paint the selection menu.
		for i, item := range _menu.order {
			if i == selected {
				screen.SetFgHex(palette.Green)
			} else {
				screen.SetFgHex(palette.Blue)
			}
			screen.Send()
			screen.MoveTo(1, y+incr)
			typewriter.WriteCenterd("[( "+item+" )]", cols)
			incr += 3
		}

		key, err := input.GetChar()
		if err != nil {
			return _menu, nil
		}
		switch key {
		case 'j', 'n':
			selected++
			if selected == len(_menu.items) {
				selected = 0
			}
		case 'k', 'p':
			selected--
			if selected < 0 {
				selected = len(_menu.items) - 1
			}
		case 'q', '\x1b':
			_menu.items["Quit"] = true
			return _menu, nil
		case '\n':
			item := _menu.order[selected]
			for k := range _menu.items {
				if k == item {
					_menu.items[k] = true
				} else {
					_menu.items[k] = false
				}
			}
			return _menu, nil
		}
	}

	/*
		screen.SaveCurPos()
		// This is not optimal, but it works...
		quitChan := make(chan int)
		var wg sync.WaitGroup
		wg.Add(1)
		go func(screen *termy.Termy) {
			defer wg.Done()
			screen.HideCur()
			defer screen.ShowCur()
			count := 0
			for {
				select {
				case <-quitChan:
					return
				default:
					os.Stdout.WriteString(".")
					time.Sleep(500 * time.Millisecond)
					count++
					if count == 5 {
						screen.RestoreCurPos()
						screen.ClearToEOL()
						count = 0
					}
				}
			}
		}(screen)

		res, err := input.GetChar()
		quitChan <- 1
		wg.Wait()
		if err != nil {
			return _menu, err
		}
	*/
}
