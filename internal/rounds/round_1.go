package rounds

import (
	"github.com/mec-nyan/kana-master/pkg/kana"
	"github.com/mec-nyan/termy"
)

func Round1Fight(screen *termy.Termy) {

	row := kana.KanaList["a"]

	roundXFight(screen, row)
}
