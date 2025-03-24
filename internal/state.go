package internal

import "github.com/mec-nyan/kana-master/pkg/kana"

type (
	Name        = string
	Description = string
)

type (
	Action        = uint
	SyllabaryMode = uint
	RoundMode     = uint

	Item struct {
		Name        string
		Value       uint
		Description string
	}

	Menu []Item
)

// Get the maximum number of columns an item.Name will occupy.
// Useful to center menu elements.
func (m Menu) MaxNameLen() int {
	mnl := 0
	for _, item := range m {
		cols := kana.CountCols(item.Name)
		if cols > mnl {
			mnl = cols
		}
	}
	return mnl
}

// Get the maximum number of columns an item.Description will occupy.
// Useful to center menu descriptions.
func (m Menu) MaxDescLen() int {
	mdl := 0
	for _, item := range m {
		cols := kana.CountCols(item.Description)
		if cols > mdl {
			mdl = cols
		}
	}
	return mdl
}

const (
	NoOp Action = iota
	Welcome
	Play
	Settings
	Back
	Select
	SelectSyllabary
	SelectGroup
	SelectRound
	Continue
	Progress
	Help
	Quit
)

const (
	_ SyllabaryMode = iota
	HiraganaMode
	KatakanaMode
	PairsMode
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
