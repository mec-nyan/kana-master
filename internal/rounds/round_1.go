package rounds

import (
	"github.com/mec-nyan/kana-master/internal"
	"github.com/mec-nyan/kana-master/pkg/kana"
	"github.com/mec-nyan/termy"
)

func Round1Fight(screen *termy.Termy, opts internal.UserOptions) {

	row := kana.List["a"]

	roundXFight(screen, row, opts)
}
