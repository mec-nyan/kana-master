package input

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mec-nyan/kana-master/pkg/kana"
)

func kanaErr(bs []byte) error {
	return fmt.Errorf("Invalid kana: %s", string(bs))
}

// TODO: What about multibyte codepoints?
func GetChar() (byte, error) {
	buff := make([]byte, 1)
	_, err := os.Stdin.Read(buff)
	return buff[0], err
}

// TODO: What about multibyte codepoints?
func GetOneOf(chars map[byte]bool) (byte, error) {
	for {
		c, err := GetChar()
		if err != nil {
			return 0, err
		}
		if chars[c] {
			return c, nil
		}
	}
}

func GetNumber(maxTries int) (int, error) {
	res := 0
	count := 0
	for {
		b, err := GetChar()
		// TODO: Printing should be handled by the caller, but we need to print one
		// character at a time so... We need to implement a function to enable "echo"
		// on the screen (see "Termy").
		os.Stdout.Write([]byte{b})
		if err != nil {
			return 0, err
		}
		// Accept
		if b == '\n' {
			return res, nil
		}
		// Abort
		if b == '\x1b' || b == 'q' {
			return -1, nil
		}
		n, err := strconv.ParseInt(string(b), 10, 0)
		if err != nil {
			return 0, err
		}
		res = res*10 + int(n)
		count++
		if count == maxTries {
			return res, nil
		}
	}
}

func GetInput() (kana.Romaji, error) {
	starts_with := map[byte]bool{
		'b': true,
		'c': true,
		'd': true,
		'f': true,
		'g': true,
		'h': true,
		'j': true,
		'k': true,
		'm': true,
		'n': true,
		'p': true,
		'r': true,
		's': true,
		't': true,
		'w': true,
		'y': true,
		'z': true,
	}

	ends_with := map[byte]bool{
		'a': true,
		'i': true,
		'u': true,
		'e': true,
		'o': true,
	}

	var romaji []byte

	first, err := GetChar()
	if err != nil {
		return "", err
	}

	os.Stdout.Write([]byte{first})

	if first == 'q' {
		return kana.Romaji(first), nil
	}

	romaji = append(romaji, first)

	// It's a vowel.
	if ends_with[first] {
		return kana.Romaji(romaji), nil
	}

	if !starts_with[first] {
		return "", kanaErr(romaji)
	}

	second, err := GetChar()
	if err != nil {
		return "", err
	}

	os.Stdout.Write([]byte{second})

	romaji = append(romaji, second)

	if starts_with[first] && ends_with[second] {
		return kana.Romaji(romaji), nil
	}

	if (second == 'h' && first == 's' || first == 'c') ||
		(second == 's' && first == 't') {
		last, err := GetChar()
		if err != nil {
			return "", err
		}
		os.Stdout.Write([]byte{last})
		romaji = append(romaji, last)
		if (last == 'i' && second == 'h') ||
			(last == 'u' && second == 's') {
			return kana.Romaji(romaji), nil
		}
	}

	return "", kanaErr(romaji)
}
