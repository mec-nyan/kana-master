package app

import (
	"os"
	"time"

	"github.com/mec-nyan/termy"
)

func MainLoop(screen *termy.Termy) error {
	welcome(screen)

	/*
	ClearScreen(screen)
	// TODO: Clearing the screen clears attributes 🤦
	screen.SetFgHex(palette.pink)
	screen.Send()

	Type("\n\tLet's start with these pairs")
	time.Sleep(1 * time.Second)
	Type("\n\tromaji: (hiragana, katakana)\n\n")
	time.Sleep(1 * time.Second)

	type kanaGroup struct {
		hg kana.Hiragana
		kk kana.Katakana
	}

	kanaList := make(map[kana.Romaji]kanaGroup, 5)
	kanaList["a"] = kanaGroup{
		hg: kana.Hg_a,
		kk: kana.Kk_a,
	}
	kanaList["i"] = kanaGroup{
		hg: kana.Hg_i,
		kk: kana.Kk_i,
	}
	kanaList["u"] = kanaGroup{
		hg: kana.Hg_u,
		kk: kana.Kk_u,
	}
	kanaList["e"] = kanaGroup{
		hg: kana.Hg_e,
		kk: kana.Kk_e,
	}
	kanaList["o"] = kanaGroup{
		hg: kana.Hg_o,
		kk: kana.Kk_o,
	}
	kanaOrder := []kana.Romaji{"a", "i", "u", "e", "o"}

	for _, k := range kanaOrder {
		v := kanaList[k]
		Type("\t" + string(k) + ": (" + string(v.hg) + ", " + string(v.kk) + ")\n")
		time.Sleep(500 * time.Millisecond)
	}

	GetChar()

	tries := 0.0
	score := 0.0
	var last kana.Romaji

Loop:
	for {
		ClearScreen(screen)
		screen.SetFg(colours.Green)
		screen.Send()

		Write("\n\tWrite in romaji:\n\n")

		// Get a random kana from the previous list.
		// Don't ask the same twice in a row.
		letter := kanaOrder[rand.Intn(len(kanaOrder))]
		if letter == last {
			continue
		}
		last = letter

		// Show the hiragana and katakana kanas.
		current := kanaList[letter]
		Write("\n\t" + string(current.hg) + " " + string(current.kk) + " => ")

		char, _ := GetChar()
		os.Stdout.Write([]byte{char})

		if char != 'q' {
			tries++
		}

		Write("\n\n\t")
		switch char {
		case 'q':
			break Loop
		case letter[0]:
			Write("Yes!!")
			score++
		default:
			Write("Nop....")
		}

		time.Sleep(time.Millisecond * 400)
	}

	ClearScreen(screen)
	screen.SetFg(colours.Green)
	screen.Send()

	Write("Good bye then!")
	perc := int(100 / tries * score)
	Write("\n\n\tYour score is: " + strconv.Itoa(perc) + "%")

	GetChar()
	*/
	ClearScreen(screen)

	screen.SetFgHex(palette.pink)
	screen.Send()

	Write("\n\tBye!")
	time.Sleep(time.Second * 1)

	screen.RestoreCurPos()
	screen.ClearToEOS()
	return nil
}

func GetChar() (byte, error) {
	buff := make([]byte, 1)
	_, err := os.Stdin.Read(buff)
	return buff[0], err
}

func ClearScreen(screen *termy.Termy) {
	screen.RestoreCurPos()
	screen.ClearToEOS()
}

func Write(str string) {
	os.Stdout.Write([]byte(str))
}

func Type(str string) {
	for _, b := range []byte(str) {
		os.Stdout.Write([]byte{b})
		time.Sleep(20 * time.Millisecond)
	}
}
