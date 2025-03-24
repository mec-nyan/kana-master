package kana

func isKana(r rune) bool {
	// This characters will be printed fullwidth in most terminals.
	// Hiragana 0x3040 .. 0x309F
	// Katakana 0x30A0 .. 0x30FF
	return r >= 0x3040 && r <= 0x30FF
}

func CountCols(str string) int {
	cols := 0
	// TODO: Account for non-printable and zero width characters.
	for _, r := range str {
		if isKana(r) {
			cols += 2
		} else {
			cols++
		}
	}
	return cols
}
