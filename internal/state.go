package internal

import "github.com/mec-nyan/kana-master/pkg/kana"

type (
	Name        = string
	Description = string
)

type (
	Action = uint

	ActionItem struct {
		Name
		Action
		Description
	}

	ActionMenu []ActionItem
)

const (
	NoOp Action = iota
	Welcome
	Play
	Settings
	Back
	SelectSyllabary
	SelectGroup
	SelectRound
	Continue
	Quit
)

type (
	SyllabaryMode = uint

	SyllabaryItem struct {
		Name
		SyllabaryMode
		Description
	}

	SyllabaryModeMenu []ModeItem
)

const (
	_ SyllabaryMode = iota
	HiraganaMode
	KatakanaMode
	PairsMode
)

type (
	RoundMode = uint

	ModeItem struct {
		Name
		RoundMode
		Description
	}

	RoundModeMenu []ModeItem
)

const (
	_ RoundMode = iota
	RowMode
	ColMode
	GroupMode
	AllMode
)

type State struct {
	Syllabary RoundMode
	Group     RoundMode
	Round     []kana.KanaRow
	Progress  struct {
		Score   float64
		Overall float64
	}
}
