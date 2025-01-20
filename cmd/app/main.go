package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mec-nyan/kana-master/pkg/kana"
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

	screen.ClearScreen()
	screen.SaveCurPos()
	screen.HideCur()
	defer screen.ShowCur()

	screen.SetFg(colours.Red)
	screen.Send()

	write("\n\t  ")

	screen.SetFgHex(palette.pink)
	screen.Send()

	write("Welcome to Kana-master!\n\n")

	screen.SetFgHex(palette.purple)
	screen.Italics()
	screen.Send()

	write("\tLearn ひらがな and カタカナ from the command line.\n\n")

	screen.SetFgHex(palette.blue)
	screen.Send()

	write("\n\tPress any key to continue...")

	getchar()
	clearScreen(screen)
	// TODO: Clearing the screen clears attributes 🤦
	screen.SetFgHex(palette.pink)
	screen.Send()

	write("\n\tLet's start with these pairs")
	write("\n\tromaji: (hiragana, katakana)\n\n")

	type kanaGroup struct {
		hg kana.Hiragana
		kk kana.Katakana
	}

	kanaList := make(map[kana.Romaji]kanaGroup, 5)
	kanaList["a"] = kanaGroup{
		hg: kana.Hg_a,
		kk: kana.Kk_a,
	}
	kanaList["i"] = kanaGroup{
		hg: kana.Hg_i,
		kk: kana.Kk_i,
	}
	kanaList["u"] = kanaGroup{
		hg: kana.Hg_u,
		kk: kana.Kk_u,
	}
	kanaList["e"] = kanaGroup{
		hg: kana.Hg_e,
		kk: kana.Kk_e,
	}
	kanaList["o"] = kanaGroup{
		hg: kana.Hg_o,
		kk: kana.Kk_o,
	}
	kanaOrder := []kana.Romaji{"a", "i", "u", "e", "o"}

	for _, k := range kanaOrder {
		v := kanaList[k]
		write("\t" + string(k) + ": (" + string(v.hg) + ", " + string(v.kk) + ")\n")
	}

	getchar()

	tries := 0.0
	score := 0.0
	var last kana.Romaji

Loop:
	for {
		clearScreen(screen)
		screen.SetFg(colours.Green)
		screen.Send()

		write("\n\tWrite in romaji:\n\n")

		// Get a random kana from the previous list.
		// Don't ask the same twice in a row.
		letter := kanaOrder[rand.Intn(len(kanaOrder))]
		if letter == last {
			continue
		}
		last = letter

		// Show the hiragana and katakana kanas.
		current := kanaList[letter]
		write("\n\t" + string(current.hg) + " " + string(current.kk) + " => ")

		char, _ := getchar()
		os.Stdout.Write([]byte{char})

		if char != 'q' {
			tries++
		}

		write("\n\n\t")
		switch char {
		case 'q':
			break Loop
		case letter[0]:
			write("Yes!!")
			score++
		default:
			write("Nop....")
		}

		time.Sleep(time.Millisecond * 400)
	}

	clearScreen(screen)
	screen.SetFg(colours.Green)
	screen.Send()

	write("Good bye then!")
	perc := int(100 / tries * score)
	write("\n\n\tYour score is: " + strconv.Itoa(perc) + "%")

	getchar()
	clearScreen(screen)

	screen.SetFgHex(palette.pink)
	screen.Send()

	write("\n\tBye!")
	time.Sleep(time.Second * 1)

	screen.RestoreCurPos()
	screen.ClearToEOS()
}

func getchar() (byte, error) {
	buff := make([]byte, 1)
	_, err := os.Stdin.Read(buff)
	return buff[0], err
}

func clearScreen(screen *termy.Termy) {
	screen.RestoreCurPos()
	screen.ClearToEOS()
}

func write(str string) {
	os.Stdout.Write([]byte(str))
}
