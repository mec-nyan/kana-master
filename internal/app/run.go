package app

import (
	"os"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/termy"
)

func Run(opts internal.CLIOptions) error {
	screen := termy.NewTermy(os.Stdout)
	err := screen.Cbreaky()
	if err != nil {
		return err
	}
	defer screen.Restore()


	err = MainLoop(screen, opts)
	if err != nil {
		return err
	}

	return nil
}
