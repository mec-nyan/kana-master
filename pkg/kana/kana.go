// Package kana associates hiragana/katakana with romaji.
package kana

type (
	// Use aliases, so we don't need to convert them later.
	// These are just for semantic purpose.
	Hiragana = string
	Katakana = string
	Romaji   = string
)

type Kana struct {
	Romaji, Alt Romaji
	Hiragana
	Katakana
}

type KanaRow []Kana

var Rows = []KanaRow{
	{
		{"a", "", Hg_a, Kk_a},
		{"i", "", Hg_i, Kk_i},
		{"u", "", Hg_u, Kk_u},
		{"e", "", Hg_e, Kk_e},
		{"o", "", Hg_o, Kk_o},
	},
	{
		{"ka", "", Hg_ka, Kk_ka},
		{"ki", "", Hg_ki, Kk_ki},
		{"ku", "", Hg_ku, Kk_ku},
		{"ke", "", Hg_ke, Kk_ke},
		{"ko", "", Hg_ko, Kk_ko},
	},
	{
		{"sa", "", Hg_sa, Kk_sa},
		{"si", "shi", Hg_si, Kk_si},
		{"su", "", Hg_su, Kk_su},
		{"se", "", Hg_se, Kk_se},
		{"so", "", Hg_so, Kk_so},
	},
	{
		{"ta", "", Hg_ta, Kk_ta},
		{"ti", "chi", Hg_ti, Kk_ti},
		{"tu", "tsu", Hg_tu, Kk_tu},
		{"te", "", Hg_te, Kk_te},
		{"to", "", Hg_to, Kk_to},
	},
	{
		{"na", "", Hg_na, Kk_na},
		{"ni", "", Hg_ni, Kk_ni},
		{"nu", "", Hg_nu, Kk_nu},
		{"ne", "", Hg_ne, Kk_ne},
		{"no", "", Hg_no, Kk_no},
	},
	{
		{"ha", "", Hg_ha, Kk_ha},
		{"hi", "", Hg_hi, Kk_hi},
		{"hu", "fu", Hg_hu, Kk_hu},
		{"he", "", Hg_he, Kk_he},
		{"ho", "", Hg_ho, Kk_ho},
	},
	{
		{"ma", "", Hg_ma, Kk_ma},
		{"mi", "", Hg_mi, Kk_mi},
		{"mu", "", Hg_mu, Kk_mu},
		{"me", "", Hg_me, Kk_me},
		{"mo", "", Hg_mo, Kk_mo},
	},
	{
		{"ya", "", Hg_ya, Kk_ya},
		{},
		{"yu", "", Hg_yu, Kk_yu},
		{},
		{"yo", "", Hg_yo, Kk_yo},
	},
	{
		{"ra", "", Hg_ra, Kk_ra},
		{"ri", "", Hg_ri, Kk_ri},
		{"ru", "", Hg_ru, Kk_ru},
		{"re", "", Hg_re, Kk_re},
		{"ro", "", Hg_ro, Kk_ro},
	},
	{
		{"wa", "", Hg_wa, Kk_wa},
		{},
		{"n ", "", Hg_n, Kk_n},
		{},
		{"wo", "", Hg_wo, Kk_wo},
	},
	{
		{"ga", "", Hg_ga, Kk_ga},
		{"gi", "", Hg_gi, Kk_gi},
		{"gu", "", Hg_gu, Kk_gu},
		{"ge", "", Hg_ge, Kk_ge},
		{"go", "", Hg_go, Kk_go},
	},
	{
		{"za", "", Hg_za, Kk_za},
		{"zi", "ji", Hg_zi, Kk_zi},
		{"zu", "", Hg_zu, Kk_zu},
		{"ze", "", Hg_ze, Kk_ze},
		{"zo", "", Hg_zo, Kk_zo},
	},
	{
		{"da", "", Hg_da, Kk_da},
		{"di", "", Hg_di, Kk_di},
		{"du", "", Hg_du, Kk_du},
		{"de", "", Hg_de, Kk_de},
		{"do", "", Hg_do, Kk_do},
	},
	{
		{"ba", "", Hg_ba, Kk_ba},
		{"bi", "", Hg_bi, Kk_bi},
		{"bu", "", Hg_bu, Kk_bu},
		{"be", "", Hg_be, Kk_be},
		{"bo", "", Hg_bo, Kk_bo},
	},
	{
		{"pa", "", Hg_pa, Kk_pa},
		{"pi", "", Hg_pi, Kk_pi},
		{"pu", "", Hg_pu, Kk_pu},
		{"pe", "", Hg_pe, Kk_pe},
		{"po", "", Hg_po, Kk_po},
	},
}

var List = map[Romaji]KanaRow{
	"a": {
		{"a", "", Hg_a, Kk_a},
		{"i", "", Hg_i, Kk_i},
		{"u", "", Hg_u, Kk_u},
		{"e", "", Hg_e, Kk_e},
		{"o", "", Hg_o, Kk_o},
	},
	"ka": {
		{"ka", "", Hg_ka, Kk_ka},
		{"ki", "", Hg_ki, Kk_ki},
		{"ku", "", Hg_ku, Kk_ku},
		{"ke", "", Hg_ke, Kk_ke},
		{"ko", "", Hg_ko, Kk_ko},
	},
	"sa": {
		{"sa", "", Hg_sa, Kk_sa},
		{"si", "shi", Hg_si, Kk_si},
		{"su", "", Hg_su, Kk_su},
		{"se", "", Hg_se, Kk_se},
		{"so", "", Hg_so, Kk_so},
	},
	"ta": {
		{"ta", "", Hg_ta, Kk_ta},
		{"ti", "chi", Hg_ti, Kk_ti},
		{"tu", "tsu", Hg_tu, Kk_tu},
		{"te", "", Hg_te, Kk_te},
		{"to", "", Hg_to, Kk_to},
	},
	"na": {
		{"na", "", Hg_na, Kk_na},
		{"ni", "", Hg_ni, Kk_ni},
		{"nu", "", Hg_nu, Kk_nu},
		{"ne", "", Hg_ne, Kk_ne},
		{"no", "", Hg_no, Kk_no},
	},
	"ha": {
		{"ha", "", Hg_ha, Kk_ha},
		{"hi", "", Hg_hi, Kk_hi},
		{"hu", "fu", Hg_hu, Kk_hu},
		{"he", "", Hg_he, Kk_he},
		{"ho", "", Hg_ho, Kk_ho},
	},
	"ma": {
		{"ma", "", Hg_ma, Kk_ma},
		{"mi", "", Hg_mi, Kk_mi},
		{"mu", "", Hg_mu, Kk_mu},
		{"me", "", Hg_me, Kk_me},
		{"mo", "", Hg_mo, Kk_mo},
	},
	"ya": {
		{"ya", "", Hg_ya, Kk_ya},
		{},
		{"yu", "", Hg_yu, Kk_yu},
		{},
		{"yo", "", Hg_yo, Kk_yo},
	},
	"ra": {
		{"ra", "", Hg_ra, Kk_ra},
		{"ri", "", Hg_ri, Kk_ri},
		{"ru", "", Hg_ru, Kk_ru},
		{"re", "", Hg_re, Kk_re},
		{"ro", "", Hg_ro, Kk_ro},
	},
	"wa": {
		{"wa", "", Hg_wa, Kk_wa},
		{},
		{"n ", "", Hg_n, Kk_n},
		{},
		{"wo", "", Hg_wo, Kk_wo},
	},
	"ga": {
		{"ga", "", Hg_ga, Kk_ga},
		{"gi", "", Hg_gi, Kk_gi},
		{"gu", "", Hg_gu, Kk_gu},
		{"ge", "", Hg_ge, Kk_ge},
		{"go", "", Hg_go, Kk_go},
	},
	"za": {
		{"za", "", Hg_za, Kk_za},
		{"zi", "ji", Hg_zi, Kk_zi},
		{"zu", "", Hg_zu, Kk_zu},
		{"ze", "", Hg_ze, Kk_ze},
		{"zo", "", Hg_zo, Kk_zo},
	},
	"da": {
		{"da", "", Hg_da, Kk_da},
		{"di", "", Hg_di, Kk_di},
		{"du", "", Hg_du, Kk_du},
		{"de", "", Hg_de, Kk_de},
		{"do", "", Hg_do, Kk_do},
	},
	"ba": {
		{"ba", "", Hg_ba, Kk_ba},
		{"bi", "", Hg_bi, Kk_bi},
		{"bu", "", Hg_bu, Kk_bu},
		{"be", "", Hg_be, Kk_be},
		{"bo", "", Hg_bo, Kk_bo},
	},
	"pa": {
		{"pa", "", Hg_pa, Kk_pa},
		{"pi", "", Hg_pi, Kk_pi},
		{"pu", "", Hg_pu, Kk_pu},
		{"pe", "", Hg_pe, Kk_pe},
		{"po", "", Hg_po, Kk_po},
	},
}

type KanaGroup struct {
	basic, dakuten, handakuten KanaRow
}

var Groups = map[Romaji]KanaGroup{
	"ka": {
		basic:   List["ka"],
		dakuten: List["ga"],
	},
	"sa": {
		basic:   List["sa"],
		dakuten: List["za"],
	},
	"ta": {
		basic:   List["ta"],
		dakuten: List["da"],
	},
	"ha": {
		basic:      List["ha"],
		dakuten:    List["ba"],
		handakuten: List["pa"],
	},
}

func GetHiragana(row KanaRow) []Hiragana {
	hiragana := []Hiragana{}
	for _, kana := range row {
		hiragana = append(hiragana, kana.Hiragana)
	}
	return hiragana
}

func GetKatakana(row KanaRow) []Katakana {
	katakana := []Katakana{}
	for _, kana := range row {
		katakana = append(katakana, kana.Katakana)
	}
	return katakana
}

func GetRomaji(row KanaRow) []Romaji {
	romaji := []Romaji{}
	for _, kana := range row {
		romaji = append(romaji, kana.Romaji)
	}
	return romaji
}
