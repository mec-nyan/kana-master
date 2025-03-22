package typewriter

import (
	"os"
	"time"
	"unicode/utf8"
)

func Write(str string) {
	os.Stdout.Write([]byte(str))
}

func Type(str string) {
	for _, b := range []byte(str) {
		os.Stdout.Write([]byte{b})
		time.Sleep(20 * time.Millisecond)
	}
}

func WriteCentered(str string, size int) {
	Write(padStr(str, size))
}

func TypeCentered(str string, size int) {
	Type(padStr(str, size))
}

func padStr(str string, size int) string {
	var padding int
	var paddedStr string
	// TODO: Count COLUMNS and not characters!
	if utf8.RuneCountInString(str) > size {
		str = str[:size]
	} else {
		padding = (size - utf8.RuneCountInString(str)) / 2
	}

	for i := 0; i < padding; i++ {
		paddedStr += " "
	}
	return string(paddedStr) + str
}
