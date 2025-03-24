// This is the entry point of our program.
package main

import (
	"flag"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/app"
)

var animate = flag.String("animation", "on", "enable/disable animations on welcome screen.")

func main() {
	flag.Parse()

	opts := internal.CLIOptions{
		Animate: *animate == "on",
	}

	err := app.Run(opts)
	if err != nil {
		panic(err)
	}
}
