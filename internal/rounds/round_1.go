package rounds

import (
	"fmt"
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

type kanaGroup struct {
	hg kana.Hiragana
	kk kana.Katakana
}

type kanaRow struct {
	list  map[kana.Romaji]kanaGroup
	order []kana.Romaji
}

var (
	goodOnes = []string{
		"Nicely done!",
		"That's it",
		"Yep!",
		"You're on fire!",
		"Well done!",
		"Cool!",
		"You've got this!",
	}
	// TODO: This must be positive comments to cheer them up!
	badOnes = []string{
		"Nop...",
		"Not even close!",
		"Really?",
		"Try again!",
		"That's not it",
		"Cat?",
		"???",
	}
)

func Round1Fight(screen *termy.Termy) {

	row := kanaRow{
		list:  makeKanaRow(),
		order: []kana.Romaji{"a", "i", "u", "e", "o"},
	}

	roundXFight(screen, row)
}

func roundXFight(screen *termy.Termy, row kanaRow) {
	intro(screen, row)

	tries, score := play(screen, row)

	end(screen, tries, score)
}

func intro(screen *termy.Termy, row kanaRow) {
	screen.ClearScreen()
	screen.Normal()
	screen.Send()

	screen.SetFgHex(palette.Blue)
	screen.Send()

	x, y := 4, 2
	screen.MoveTo(x, y)
	typewriter.Type("Let's start with these pairs")
	time.Sleep(500 * time.Millisecond)
	y += 2
	screen.MoveTo(x, y)
	typewriter.Type("romaji: (hiragana, katakana)")
	time.Sleep(500 * time.Millisecond)

	screen.SetFg(colours.White)
	screen.Send()
	for _, k := range row.order {
		y += 2
		screen.MoveTo(x, y)
		v := row.list[k]
		typewriter.Type(string(k) + ": (" + string(v.hg) + ", " + string(v.kk) + ")")
		time.Sleep(500 * time.Millisecond)
	}

	screen.SetFgHex(palette.Blue)
	screen.Send()

	y += 4
	screen.MoveTo(x, y)
	typewriter.Type("Press any key to start!")

	input.GetChar()
}

func makeKanaRow() map[kana.Romaji]kanaGroup {
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
	return kanaList
}

func play(screen *termy.Termy, row kanaRow) (float64, float64) {
	tries := 0.0
	score := 0.0
	progress := map[kana.Romaji]int{}
	var last kana.Romaji

	screen.ClearScreen()
	screen.SetFg(colours.Blue)
	screen.Send()

	x, y := 4, 2
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
		x = 8

		screen.RestoreCurPos()
		screen.ClearToEOS()

		screen.MoveTo(x, y+10)
		screen.ClearToEOL()
		screen.SetFgHex(palette.Purple)
		screen.Send()

		fmt.Printf("<Progress> a %2d -- i %2d -- u %2d -- e %2d -- o %2d",
			progress["a"], progress["i"], progress["u"], progress["e"], progress["o"])

		screen.SetFgHex(palette.Green)
		screen.Send()

		// Get a random kana from the previous list.
		// Don't ask the same twice in a row.
		letter := row.order[rand.Intn(len(row.order))]
		if letter == last {
			continue
		}
		last = letter

		// Show the hiragana and katakana kanas.
		current := row.list[letter]
		y += 2
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
			// TODO: Add different messages.
			typewriter.Write(goodOnes[rand.Intn(len(goodOnes))])
			progress[letter]++
			score++
		default:
			// TODO: What to do if the user misses?
			if progress[letter] > 0 {
			}
			typewriter.Write(badOnes[rand.Intn(len(badOnes))])
		}

		time.Sleep(600 * time.Millisecond)
	}

	return tries, score
}

func end(screen *termy.Termy, tries, score float64) {
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
