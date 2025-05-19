// This is the entry point of our program.
package main

import (
	"flag"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/app"
)

var noAnim, palette bool

func init() {
	flag.BoolVar(&noAnim, "no-anim", false, "Disable animations.")
	flag.BoolVar(&noAnim, "na", false, "Disable animations (shorthand).")
	// TODO: The logic to switch between the terminal theme and a custom palette.
	flag.BoolVar(&palette, "palette", false, "Use custom palette.")
	flag.BoolVar(&palette, "p", false, "Use custom palette (shorthand).")
}

func main() {

	flag.Parse()

	opts := internal.UserOptions{
		Animate:       !noAnim,
		CustomPalette: palette,
		// Unless otherwise selected, practice with pairs.
		PractisePairs: true,
	}

	err := app.Run(opts)
	if err != nil {
		panic(err)
	}
}
