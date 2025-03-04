package main


import "github.com/mec-nyan/kana-master/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
