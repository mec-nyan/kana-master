package main

import (
	"log"
	"os"
	"time"

	"github.com/mec-nyan/termy"
	"github.com/mec-nyan/termy/colours"
)

var palette = struct{ purple, pink, blue string }{purple: "CA66F4", pink: "F466D8", blue: "8366f4"}

func main() {
	fd := int(os.Stdout.Fd())
	term := termy.New(fd, false)
	err := term.Cbreaky()
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer term.Restore()

	screen := termy.NewTermy(os.Stdout)

	screen.SaveCurPos()
	screen.HideCur()
	defer screen.ShowCur()

	screen.SetFg(colours.Red)
	screen.Send()

	os.Stdout.Write([]byte("   "))

	screen.SetFgHex(palette.pink)
	screen.Send()

	os.Stdout.Write([]byte("Welcome to Kana-master!\n\n"))

	screen.SetFgHex(palette.purple)
	screen.Italics()
	screen.Send()

	os.Stdout.Write([]byte("\tLearn ひらがな and カタカナ from the command line.\n\n"))

	screen.SetFgHex(palette.blue)
	screen.Send()

	os.Stdout.Write([]byte("\tPress any key to continue..."))

	buff := make([]byte, 1)
	os.Stdin.Read(buff)

	screen.RestoreCurPos()
	screen.ClearToEOS()

	screen.SetFgHex(palette.pink)
	screen.Send()

	os.Stdout.Write([]byte("\tBye for now!"))
	time.Sleep(time.Second * 1)

	screen.RestoreCurPos()
	screen.ClearToEOS()
}
