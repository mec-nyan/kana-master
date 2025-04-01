package app

import (
	"os"

	"github.com/mec-nyan/termy"

	"github.com/mec-nyan/kana-master/internal"
)

func Run(opts internal.UserOptions) error {
	display, err := termy.NewDisplay(os.Stdout)
	if err != nil {
		return err
	}

	// Disable stdin buffering.
	err = display.UnCookIt()
	if err != nil {
		return err
	}

	// Don't echo user input.
	err = display.NoEcho()
	if err != nil {
		return err
	}

	defer display.Restore()

	err = MainLoop(display, opts)
	if err != nil {
		return err
	}

	return nil
}
