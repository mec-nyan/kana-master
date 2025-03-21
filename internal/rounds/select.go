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

	putCenteredAt(screen, "Round Selection", 2, writeFunc)

	rows, _, _ := screen.Size()
	xPos, yPos := 4, 4

	index := 1
	for _, row := range kana.Rows {
		if yPos > rows-4 {
			yPos = 4
			// TODO: Check against the width of the screen!
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
		return kana.KanaRow{}, "quit"
	}

	return kana.Rows[i-1], "play"
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

func NextRow(row kana.KanaRow) kana.KanaRow {
	current := row[0].Romaji
	rowIdx := 0
	for i, r := range kana.Rows {
		if r[0].Romaji == current {
			rowIdx = i + 1
			break
		}
	}
	if rowIdx == len(kana.Rows) {
		// Wrap around!
		// TODO: Say you reach the end of the game/stage.
		rowIdx = 0
	}
	return kana.Rows[rowIdx]
}
