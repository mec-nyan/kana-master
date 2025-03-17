package rounds

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mec-nyan/kana-master/internal/input"
	"github.com/mec-nyan/kana-master/internal/palette"
	"github.com/mec-nyan/kana-master/internal/typewriter"
	"github.com/mec-nyan/kana-master/pkg/kana"
	"github.com/mec-nyan/termy"
	"github.com/mec-nyan/termy/colours"
)

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

func roundXFight(screen *termy.Termy, row kana.KanaRow) {
	intro(screen, row)

	tries, score := play(screen, row)

	end(screen, tries, score)
}

func intro(screen *termy.Termy, row kana.KanaRow) {
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
	for _, v := range row {
		y += 2
		screen.MoveTo(x, y)
		typewriter.Type(string(v.Romaji) + ": (" + string(v.Hiragana) + ", " + string(v.Katakana) + ")")
		time.Sleep(500 * time.Millisecond)
	}

	screen.SetFgHex(palette.Blue)
	screen.Send()

	y += 4
	screen.MoveTo(x, y)
	typewriter.Type("Press any key to start!")

	input.GetChar()
}

// play will play the game as follows:
// Take the row consisting of five kanas (one for each vowel in "a i u e o"),
// Shuffle,
// play the row,
// Shuffle again makeing sure we don't start with the same kana
// we ended the las row,
// Keep going until the goal (i.e. five correct answers for each kana) is reached.
func play(screen *termy.Termy, row kana.KanaRow) (float64, float64) {
	// We'll use rand to shuffle the rows each time.
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Make a copy of the row.
	prefShuffle := row[:]

	// We'll keep track of the tries and the score to know the overall rating.
	tries := 0.0
	score := 0.0
	// And show some progress (WIP)
	progress := map[kana.Romaji]int{}

	// We'll redraw this every frame.
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
		// Shuffle the kanas on every turn.
		shuffledRow := prefShuffle[:]
		rnd.Shuffle(len(shuffledRow), func(i, j int) {
			shuffledRow[i], shuffledRow[j] = shuffledRow[j], shuffledRow[i]
		})
		// Make shure it doesn't start where we ended last time.
		if shuffledRow[0] == prefShuffle[len(prefShuffle)-1] {
			continue
		}
		// Play the row.
		for _, currentKana := range shuffledRow {
			y := y
			x = 8

			screen.RestoreCurPos()
			screen.ClearToEOS()

			screen.MoveTo(x, y+10)
			screen.ClearToEOL()
			screen.SetFgHex(palette.Blue)
			screen.Send()

			// TODO: Nice progress bars.
			fmt.Printf("Progress:")
			screen.SetFg(1)
			screen.Send()
			screen.MoveTo(x, y+11)
			screen.ClearToEOL()
			fmt.Printf("a  (%-5s)", strings.Repeat("▄", progress["a"]))
			screen.SetFgHex(palette.Green)
			screen.Send()
			screen.MoveTo(x, y+12)
			screen.ClearToEOL()
			fmt.Printf("i  (%-5s)", strings.Repeat("▄", progress["i"]))
			screen.SetFgHex(palette.Pink)
			screen.Send()
			screen.MoveTo(x, y+13)
			screen.ClearToEOL()
			fmt.Printf("u  (%-5s)", strings.Repeat("▄", progress["u"]))
			screen.SetFgHex(palette.Blue)
			screen.Send()
			screen.MoveTo(x, y+14)
			screen.ClearToEOL()
			fmt.Printf("e  (%-5s)", strings.Repeat("▄", progress["e"]))
			screen.SetFgHex(palette.Purple)
			screen.Send()
			screen.MoveTo(x, y+15)
			screen.ClearToEOL()
			fmt.Printf("o  (%-5s)", strings.Repeat("▄", progress["o"]))

			screen.SetFgHex(palette.Green)
			screen.Send()

			// Show the hiragana and katakana kanas.
			y += 2
			screen.MoveTo(x, y)
			typewriter.Write(string(currentKana.Hiragana) + " " + string(currentKana.Katakana) + " => ")

			// TODO: This works for "a, i, u, e, o".
			// How do we make it work for i.e. "ka, ki, ku, ..."?
			char, _ := input.GetChar()
			os.Stdout.Write([]byte{char})

			if char != 'q' {
				tries++
			}

			letter := currentKana.Romaji

			y += 2
			screen.MoveTo(x, y)

			switch char {
			case 'q', '\x1b':
				break Loop
			case letter[0]:
				// TODO: Add different messages.
				typewriter.Write(goodOnes[rand.Intn(len(goodOnes))])
				if progress[letter] < 5 {
					progress[letter]++
				}
				score++
			default:
				// TODO: What to do if the user misses?
				if progress[letter] > 0 {
					progress[letter]--
				}
				typewriter.Write(badOnes[rand.Intn(len(badOnes))])
			}
			time.Sleep(600 * time.Millisecond)

			if progress["a"] >= 5 && progress["i"] >= 5 && progress["u"] >= 5 &&
				progress["e"] >= 5 && progress["o"] >= 5 {
				break Loop
			}
		}
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
