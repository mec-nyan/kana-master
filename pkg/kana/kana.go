// Package kana associates hiragana/katakana with romaji.
package kana

type (
	Hiragana rune
	Katakana rune
	Romaji   string
)

const (
	hg_a Hiragana = 'あ'
	hg_i Hiragana = 'い'
	hg_u Hiragana = 'う'
	hg_e Hiragana = 'え'
	hg_o Hiragana = 'お'
	// ...

	kk_a Katakana = 'ア'
	kk_i Katakana = 'イ'
	kk_u Katakana = 'ウ'
	kk_e Katakana = 'エ'
	kk_o Katakana = 'オ'
	// ...
)

type Kana struct {
	Hiragana
	Katakana
	Romaji
}

var (
	hiragana map[Romaji]Hiragana
	katakana map[Romaji]Katakana
	kana     map[Romaji]Kana
)
