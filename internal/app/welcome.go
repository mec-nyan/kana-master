package app

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

const header = `
█   █                      █▀▀█▀▀█                               
█   █  █████ █████ █████   █  █  █ █████ █████ █████ ████ █████  
██████ █   █ █   █ █   █   ██ █  █ █   █ █   ▀   █   █    █   █  
██   █ █████ ██  █ █████   ██ █  █ █████ █████   ██  ████ ██████ 
██   █ ██  █ ██  █ ██  █   ██ █  █ ██  █    ██   ██  ██   ██   █ 
██   █ ██  █ ██  █ ██  █   ██ █  █ ██  █ █████   ██  ████ ██   █ 
`

func welcome(screen *termy.Termy, animate bool) {
	printFunc := typewriter.Write
	delay := 0 * time.Millisecond
	if animate {
		printFunc = typewriter.Type
		delay = 500 * time.Millisecond
	}

	_, cols, _ := screen.Size()

	screen.ClearScreen()
	screen.SaveCurPos()
	screen.SetFgHex(palette.Grey)
	screen.Send()

	typewriter.Write("Mec-Nyan's Kana-Master v0.1.0-beta")
	screen.CurToCol(1)
	screen.MoveDown(1)
	typewriter.Write("License information and stuff will be here.")

	headerLines := strings.Split(header, "\n")
	headerWidth := utf8.RuneCountInString(headerLines[1])

	headerX := (cols - headerWidth) / 2
	headerY := 4

	screen.SetFgHex(palette.Blue)
	screen.Send()

	for _, line := range headerLines {
		screen.MoveTo(headerX, headerY)
		typewriter.Write(line)
		headerY++
	}

	screen.SetFgHex(palette.Purple)
	screen.Send()

	y := headerY + 4
	screen.MoveTo(8, y)
	printFunc("  Welcome to Kana-master!")
	time.Sleep(delay)

	y += 2
	screen.MoveTo(8, y)
	printFunc("Learn ひらがな and カタカナ from the command line.")
	time.Sleep(delay)

	screen.SetFgHex(palette.Blue)
	screen.Italics()
	screen.Send()

	y += 6
	screen.MoveTo(8, y)
	printFunc("To continue, press any key")

	input.GetChar()
}
