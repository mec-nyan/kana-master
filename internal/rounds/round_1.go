package rounds

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/kana-master/pkg/kana"
	"github.com/mec-nyan/termy"
	"github.com/mec-nyan/termy/colours"
)

func Round1Fight(screen *termy.Termy, _ *termy.TermSettings) {
	screen.ClearScreen()
	screen.Normal()
	screen.Send()

	screen.SetFgHex(palette.Blue)
	screen.Send()

	x, y := 4, 2
	screen.MoveTo(x, y)
	typewriter.Type("Let's start with these pairs")
	time.Sleep(1 * time.Second)
	y += 2
	screen.MoveTo(x, y)
	typewriter.Type("romaji: (hiragana, katakana)")
	time.Sleep(1 * time.Second)

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

	screen.SetFg(colours.White)
	screen.Send()
	for _, k := range kanaOrder {
		y += 2
		screen.MoveTo(x, y)
		v := kanaList[k]
		typewriter.Type(string(k) + ": (" + string(v.hg) + ", " + string(v.kk) + ")")
		time.Sleep(500 * time.Millisecond)
	}

	screen.SetFgHex(palette.Blue)
	screen.Send()

	y += 4
	screen.MoveTo(x, y)
	typewriter.Type("Press any key to start!")

	input.GetChar()

	tries := 0.0
	score := 0.0
	var last kana.Romaji

	screen.ClearScreen()
	screen.SetFg(colours.Blue)
	screen.Send()

	y = 2
	screen.MoveTo(x, y)
	typewriter.Write("Write in romaji:")
	screen.SetFgHex(palette.Grey)
	screen.Send()
	y += 1
	screen.MoveTo(x, y)
	typewriter.Write("(press \"q\" to end this round)")
	screen.SaveCurPos()

Loop:
	for {
		y := y
		screen.RestoreCurPos()
		screen.ClearToEOS()

		screen.SetFgHex(palette.Green)
		screen.Send()

		// Get a random kana from the previous list.
		// Don't ask the same twice in a row.
		letter := kanaOrder[rand.Intn(len(kanaOrder))]
		if letter == last {
			continue
		}
		last = letter

		// Show the hiragana and katakana kanas.
		current := kanaList[letter]
		y += 2
		x = 8
		screen.MoveTo(x, y)
		typewriter.Write(string(current.hg) + " " + string(current.kk) + " => ")

		char, _ := input.GetChar()
		os.Stdout.Write([]byte{char})

		if char != 'q' {
			tries++
		}

		y += 2
		screen.MoveTo(x, y)
		switch char {
		case 'q':
			break Loop
		case letter[0]:
			typewriter.Write("Good job!")
			score++
		default:
			typewriter.Write("Not even close!")
		}

		time.Sleep(time.Millisecond * 400)
	}

	screen.ClearScreen()
	screen.SetFg(colours.Yellow)
	screen.Send()

	screen.MoveTo(4, 2)
	perc := int(100 / tries * score)
	typewriter.Type("You've scored " + strconv.Itoa(perc) + "%")

	screen.MoveTo(4, 5)
	typewriter.Write("(press any key)")

	input.GetChar()

}
