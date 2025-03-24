package rounds

import (
	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/pkg/kana"
)

type Kind uint

const (
	Row Kind = iota
	Column
	Group
	All
)


type Round struct {
	Name string
	// Set can represent a row, a group, a column, or all.
	// I.e.:
	// row: { あ, い, う, え, お, }
	// column: { あ, か, さ, た, な, は, 。。。 }
	// group: { ( か, き, く, け, こ ), ( が, ぎ, ぐ, げ, ご ) }
	// all: All hiragana and/or katakana.
	// NOTE: We don't include compounds yet like きゃ etc.
	// NOTE: When selecting a column or all, should we give the
	// option to include or exclude voiced variants?
	// TODO: Use the proper terminology (五十音、よおん)
	Set []kana.Kana
	Kind
}

func NewRound(name string, set []kana.Kana, kind Kind) Round {
	return Round{
		Name: name,
		Set: set,
		Kind: kind,
	}
}

func NextRow(row kana.KanaRow) kana.KanaRow {
	current := row[0].Romaji
	rowIdx := 0
	for i, r := range kana.Rows {
		if r[0].Romaji == current {
			rowIdx = i + 1
			break
		}
	}
	if rowIdx == len(kana.Rows) {
		// Wrap around!
		// TODO: Say you reach the end of the game/stage.
		rowIdx = 0
	}
	return kana.Rows[rowIdx]
}

type RoundMode struct {
	Syllabary internal.SyllabaryMode
	Group     internal.RoundMode
}

func getRounds(mode internal.RoundMode) []kana.KanaRow {
	var rounds []kana.KanaRow
	switch mode {
	case internal.RowMode:
		for _, row := range kana.Rows {
			rounds = append(rounds, row)
		}
	case internal.ColMode:
		cols := [5]kana.KanaRow{}
		for _, row := range kana.Rows {
			for i, k := range row {
				cols[i] = append(cols[i], k)
			}
		}
		for _, col := range cols {
			rounds = append(rounds, col)
		}
		// TODO: How to separate regular from dakuten.
		// No need to return anything for "all".
	}
	return rounds
}
