package quit

import (
	"errors"
	"os"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/termy"
)

const (
	key_enter  = '\x0a'
	key_escape = '\x1b'
)

func Confirm(display *termy.Display, previous internal.Action) (internal.Action, error) {
	display.ClearScreen()

	rows, cols, err := display.Size()
	if err != nil {
		return internal.NoOp, err
	}

	message := []string{
		"Press <ESC> to cancel and go back.",
		"Press <ENTER> to accept and  quit.",
	}

	x_pos := (cols - len(message[0])) / 2
	y_pos := 4

	display.SetFg(8).Italics(true).Send()
	display.MoveTo(x_pos, y_pos)
	os.Stdout.WriteString(message[0])
	y_pos++
	display.MoveTo(x_pos, y_pos)
	os.Stdout.WriteString(message[1])

	prompt := "Are you sure? "

	x_pos = (cols - len(prompt)) / 2
	y_pos = rows / 2

	display.Normal().SetFg(2).Send()
	display.MoveTo(x_pos, y_pos)
	os.Stdout.WriteString(prompt)

	answer, err := input.GetOneOf(map[byte]bool{
		key_enter:  true,
		key_escape: true,
	})

	if err != nil {
		return internal.NoOp, err
	}

	switch answer {
	case key_enter:
		return internal.Exit, nil
	case key_escape:
		return previous, nil
	}

	return internal.NoOp, errors.New("failed to get user input!")
}
