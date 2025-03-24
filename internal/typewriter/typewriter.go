package typewriter

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mec-nyan/kana-master/pkg/kana"
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
	output, _ := PadStr(str, size)
	Write(output)
}

func TypeCentered(str string, size int) {
	output, _ := PadStr(str, size)
	Type(output)
}

func PadStr(str string, size int) (string, error) {
	var padding int
	var paddedStr string
	// TODO: Count COLUMNS and not characters!
	cols := kana.CountCols(str)
	if cols > size {
		return "", fmt.Errorf("err: str \"%s\" occupies more than %d cols", str, size)
	} else {
		padding = (size - cols) / 2
	}

	for i := 0; i < padding; i++ {
		paddedStr += " "
	}
	return string(paddedStr) + str, nil
}

func CenterStr(str string, width int) (string, error) {
	cols := kana.CountCols(str)
	if cols > width {
		return "", fmt.Errorf("err: str \"%s\" occupies more than %d cols", str, width)
	}
	padding := width - len(str)
	leftPad := padding / 2
	rightPad := padding - leftPad

	return strings.Repeat(" ", leftPad) + str + strings.Repeat(" ", rightPad), nil
}
