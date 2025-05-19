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

func Fight(display *termy.Display, row kana.KanaRow, opts internal.UserOptions) (internal.Action, error) {
	intro(display, row, opts)

	tries, score := play(display, row, opts)

	return end(display, tries, score, opts)
}

func intro(display *termy.Display, row kana.KanaRow, opts internal.UserOptions) {
	write := typewriter.Write
	delay := 0 * time.Millisecond
	if opts.Animate {
		delay = 500 * time.Millisecond
		write = typewriter.Type
	}
	display.ClearScreen()
	display.Normal()
	display.Send()

	display.SetFgHex(palette.Blue)
	display.Send()

	x, y := 4, 2
	display.MoveTo(x, y)
	write("Let's start with these sounds.")
	time.Sleep(delay)
	y += 2
	display.MoveTo(x, y)
	if opts.PractisePairs {
		write("romaji - (hiragana, katakana)")
	} else if opts.PractiseHiragana {
		write("romaji - hiragana")
	} else if opts.PractiseKatakana {
		write("romaji - katakana")
	}
	time.Sleep(delay)

	display.SetFg(colours.White)
	display.Send()
	for _, v := range row {
		y += 2
		display.MoveTo(x, y)
		if opts.PractisePairs {
			write(string(v.Romaji) + " - (" + string(v.Hiragana) + ", " + string(v.Katakana) + ")")
		} else if opts.PractiseHiragana {
			write(string(v.Romaji) + " - " + string(v.Hiragana))
		} else if opts.PractiseKatakana {
			write(string(v.Romaji) + " - " + string(v.Katakana))
		}
		time.Sleep(delay)
	}

	display.SetFgHex(palette.Blue)
	display.Send()

	y += 4
	display.MoveTo(x, y)
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
func play(display *termy.Display, row kana.KanaRow, opts internal.UserOptions) (float64, float64) {
	write := typewriter.Write
	if opts.Animate {
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
	display.ClearScreen()
	display.SetFg(colours.Blue)
	display.Send()

	x, y := 4, 2
	display.MoveTo(x, y)
	write("Write in romaji:")
	display.SetFgHex(palette.Grey)
	display.Send()
	y += 1
	display.MoveTo(x, y)
	write("(press \"q\" to end this round)")
	display.SaveCurPos()

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
		// Save current shuffle.
		prevShuffle = shuffledRow[:]
		// Play the row.
		for _, currentKana := range shuffledRow {
			y := y
			x = 8

			display.RestoreCurPos()
			display.ClearToEOS()

			// Show progress bars.
			display.MoveTo(x, y+10)
			display.ClearToEOL()
			display.SetFgHex(palette.Blue)
			display.Send()

			fmt.Printf("Progress:")
			for i, v := range order {
				display.SetFg(6 - i)
				display.Send()
				display.MoveTo(x, y+11+i)
				display.ClearToEOL()
				fmt.Printf("%s  (%-5s)", v, strings.Repeat("▄", progress[v]))
			}

			display.SetFgHex(palette.Green)
			display.Send()

			// Show the hiragana and katakana kanas.
			y += 2
			display.MoveTo(x, y)
			if opts.PractisePairs {
				write(string(currentKana.Hiragana) + " " + string(currentKana.Katakana) + " => ")
			} else if opts.PractiseHiragana {
				write(string(currentKana.Hiragana) + " => ")
			} else if opts.PractiseKatakana {
				write(string(currentKana.Katakana) + " => ")
			}

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
			display.MoveTo(x, y)

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

func end(display *termy.Display, tries, score float64, opts internal.UserOptions) (internal.Action, error) {
	write := typewriter.Write
	if opts.Animate {
		write = typewriter.Type
	}
	display.ClearScreen()
	display.SetFg(colours.Yellow)
	display.Send()

	display.MoveTo(4, 2)
	perc := int(100 / tries * score)
	write("You've scored " + strconv.Itoa(perc) + "%")

	display.MoveTo(4, 5)
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
