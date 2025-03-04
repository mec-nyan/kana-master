package input

import "os"

func GetChar() (byte, error) {
	buff := make([]byte, 1)
	_, err := os.Stdin.Read(buff)
	return buff[0], err
}
