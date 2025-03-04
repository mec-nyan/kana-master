package app

import (
	"github.com/mec-nyan/termy"
	"os"
)

func Run() error {
	fd := int(os.Stdout.Fd())
	term := termy.New(fd, false)
	err := term.Cbreaky()
	if err != nil {
		return err
	}
	defer term.Restore()

	screen := termy.NewTermy(os.Stdout)

	err = MainLoop(screen, term)
	if err != nil {
		return err
	}

	return nil
}
