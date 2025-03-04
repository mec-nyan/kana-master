package typewriter

import (
	"os"
	"time"
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
