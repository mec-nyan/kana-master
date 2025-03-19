package main

import (
	"flag"

	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/internal/app"
)

var animate = flag.String("animation", "on", "enable/disable animations on welcome screen.")
var showWelcomeScreen = flag.Bool("welcome", true, "show welcome screen.")
var showOptionsScreen = flag.Bool("options", true, "show options screen.")

func main() {
	flag.Parse()

	opts := internal.CLIOptions{
		Animate: *animate == "on",
		ShowWelcomeScreen: *showWelcomeScreen,
		ShowOptionsScreen: *showOptionsScreen,
	}

	err := app.Run(opts)
	if err != nil {
		panic(err)
	}
}
