package internal

type (
	CLIOptions struct {
		Animate           bool
		ShowWelcomeScreen bool
		ShowOptionsScreen bool
	}

	UserOptions struct {
		AnimationOn      bool
		UsePalette       bool
		PractisePairs    bool
		PractiseHiragana bool
		PractiseKatakana bool
		BackToMain       bool
	}
)
