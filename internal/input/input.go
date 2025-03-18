package input

import (
	"fmt"
	"os"

	"github.com/mec-nyan/kana-master/pkg/kana"
)

func kanaErr(bs []byte) error {
	return fmt.Errorf("Invalid kana: %s", string(bs))
}

func GetChar() (byte, error) {
	buff := make([]byte, 1)
	_, err := os.Stdin.Read(buff)
	return buff[0], err
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

	os.Stdout.Write([]byte{first})

	romaji = append(romaji, first)

	if starts_with[first] && ends_with[second] {
		return kana.Romaji(romaji), nil
	}

	if (second == 'h' && first == 's' || first == 'c') ||
		(second == 's' && first == 't') {
		last, err := GetChar()
		if err != nil {
			return "", err
		}
		romaji = append(romaji, last)
		if (last == 'i' && second == 'h') ||
			(last == 'u' && second == 's') {
			return kana.Romaji(romaji), nil
		}
	}

	return "", kanaErr(romaji)
}
