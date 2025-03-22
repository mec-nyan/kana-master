package play

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/mec-nyan/kana-master/internal"
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

func Fight(screen *termy.Termy, row kana.KanaRow, opts internal.UserOptions) (internal.Action, error) {
	intro(screen, row, opts)

	tries, score := play(screen, row, opts)

	return end(screen, tries, score, opts)
}

func intro(screen *termy.Termy, row kana.KanaRow, opts internal.UserOptions) {
	write := typewriter.Write
	delay := 0 * time.Millisecond
	if opts.AnimationOn {
		delay = 500 * time.Millisecond
		write = typewriter.Type
	}
	screen.ClearScreen()
	screen.Normal()
	screen.Send()

	screen.SetFgHex(palette.Blue)
	screen.Send()

	x, y := 4, 2
	screen.MoveTo(x, y)
	write("Let's start with these pairs")
	time.Sleep(delay)
	y += 2
	screen.MoveTo(x, y)
	write("romaji: (hiragana, katakana)")
	time.Sleep(delay)

	screen.SetFg(colours.White)
	screen.Send()
	for _, v := range row {
		y += 2
		screen.MoveTo(x, y)
		write(string(v.Romaji) + ": (" + string(v.Hiragana) + ", " + string(v.Katakana) + ")")
		time.Sleep(delay)
	}

	screen.SetFgHex(palette.Blue)
	screen.Send()

	y += 4
	screen.MoveTo(x, y)
	write("Press any key to start!")

	input.GetChar()
}

// play will play the game as follows:
// Take the row consisting of five kanas (one for each vowel in "a i u e o"),
// Shuffle,
// play the row,
// Shuffle again makeing sure we don't start with the same kana
// we ended the las row,
// Keep going until the goal (i.e. five correct answers for each kana) is reached.
func play(screen *termy.Termy, row kana.KanaRow, opts internal.UserOptions) (float64, float64) {
	write := typewriter.Write
	if opts.AnimationOn {
		write = typewriter.Type
	}
	// We'll use rand to shuffle the rows each time.
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Make a copy of the row.
	prevShuffle := row[:]

	var order []kana.Romaji
	for _, k := range row {
		order = append(order, k.Romaji)
	}
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
	write("Write in romaji:")
	screen.SetFgHex(palette.Grey)
	screen.Send()
	y += 1
	screen.MoveTo(x, y)
	write("(press \"q\" to end this round)")
	screen.SaveCurPos()

Loop:
	for {
		// Shuffle the kanas on every turn.
		shuffledRow := prevShuffle[:]
		rnd.Shuffle(len(shuffledRow), func(i, j int) {
			shuffledRow[i], shuffledRow[j] = shuffledRow[j], shuffledRow[i]
		})
		// Make shure it doesn't start where we ended last time.
		if shuffledRow[0] == prevShuffle[len(prevShuffle)-1] {
			continue
		}
		// Play the row.
		for _, currentKana := range shuffledRow {
			y := y
			x = 8

			screen.RestoreCurPos()
			screen.ClearToEOS()

			// Show progress bars.
			screen.MoveTo(x, y+10)
			screen.ClearToEOL()
			screen.SetFgHex(palette.Blue)
			screen.Send()

			fmt.Printf("Progress:")
			for i, v := range order {
				screen.SetFg(6 - i)
				screen.Send()
				screen.MoveTo(x, y+11+i)
				screen.ClearToEOL()
				fmt.Printf("%s  (%-5s)", v, strings.Repeat("▄", progress[v]))
			}

			screen.SetFgHex(palette.Green)
			screen.Send()

			// Show the hiragana and katakana kanas.
			y += 2
			screen.MoveTo(x, y)
			write(string(currentKana.Hiragana) + " " + string(currentKana.Katakana) + " => ")

			// TODO: Handle error here.
			kana, _ := input.GetInput()

			if kana != "q" && kana != "\x1b" {
				tries++
			}

			// Use double "n" to insert ん
			if kana == "nn" {
				kana = "n"
			}

			letter := currentKana.Romaji
			letterAlt := currentKana.Alt

			y += 2
			screen.MoveTo(x, y)

			switch kana {
			case "q", "\x1b":
				break Loop
			case letter, letterAlt:
				// TODO: Add different messages.
				write(goodOnes[rand.Intn(len(goodOnes))])
				if progress[letter] < 5 {
					progress[letter]++
				}
				score++
			default:
				// TODO: What to do if the user misses?
				if progress[letter] > 0 {
					progress[letter]--
				}
				write(badOnes[rand.Intn(len(badOnes))])
			}
			time.Sleep(800 * time.Millisecond)

			finish := true
			for _, prog := range progress {
				if prog < 5 {
					finish = false
					break
				}
			}
			if finish {
				break Loop
			}
		}
	}
	return tries, score
}

func end(screen *termy.Termy, tries, score float64, opts internal.UserOptions) (internal.Action, error) {
	write := typewriter.Write
	if opts.AnimationOn {
		write = typewriter.Type
	}
	screen.ClearScreen()
	screen.SetFg(colours.Yellow)
	screen.Send()

	screen.MoveTo(4, 2)
	perc := int(100 / tries * score)
	write("You've scored " + strconv.Itoa(perc) + "%")

	screen.MoveTo(4, 5)
	write("Press any key to continue (q quits)")

	res, err := input.GetChar()
	if err != nil {
		return internal.NoOp, err
	}
	switch res {
	case 'q':
		return internal.Quit, nil
	default:
		return internal.Continue, nil
	}
}
