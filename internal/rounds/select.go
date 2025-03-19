package rounds

import (
	"errors"
	"unicode/utf8"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/termy"
)

// SelectRound presents a screen with information about completed
// rounds, score, overall, and lets the user select where to go
// from here.
func SelectRound(screen *termy.Termy, opts internal.UserOptions) (
	round int, quit bool,
) {
	writeFunc := typewriter.Write
	if opts.AnimationOn {
		writeFunc = typewriter.Type
	}

	screen.ClearScreen()
	screen.UseDefault()
	screen.Send()

	putCenteredAt(screen, "Round Selection", 4, writeFunc)
	putCenteredAt(screen, "press any key", 14, writeFunc)

	input.GetChar()

	return 0, true
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
