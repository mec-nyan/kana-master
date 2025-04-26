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

// We'll group the kanas in several ways, for practice.
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
	Monographs, Digraphs KanaRow
}

// Let's try to group kanas by their "main" consonant.
// Note: Not every row will have diacitics nor digraphs.
// I.e. the "vowels" row, etc.

type KanaSet struct {
	Basic, Dakuten, Handakuten KanaGroup
}

type KanaTable map[string]KanaSet

var Table = KanaTable{
	"a": {
		Basic: KanaGroup{
			Monographs: List["a"],
		},
	},
	"k": {
		Basic: KanaGroup{
			Monographs: List["ka"],
			Digraphs: KanaRow{
				{"kya", "", Hg_kya, Kk_kya},
				{"kyu", "", Hg_kyu, Kk_kyu},
				{"kyo", "", Hg_kyo, Kk_kyo},
			},
		},
		Dakuten: KanaGroup{
			Monographs: List["ga"],
			Digraphs: KanaRow{
				{"gya", "", Hg_gya, Kk_gya},
				{"gyu", "", Hg_gyu, Kk_gyu},
				{"gyo", "", Hg_gyo, Kk_gyo},
			},
		},
	},
	"s": {
		Basic: KanaGroup{
			Monographs: List["sa"],
			Digraphs: KanaRow{
				{"sya", "sha", Hg_sya, Kk_sya},
				{"syu", "shu", Hg_syu, Kk_syu},
				{"syo", "sho", Hg_syo, Kk_syo},
			},
		},
		Dakuten: KanaGroup{
			Monographs: List["za"],
			Digraphs: KanaRow{
				{"zya", "ja", Hg_zya, Kk_zya},
				{"zyu", "ju", Hg_zyu, Kk_zyu},
				{"zyo", "jo", Hg_zyo, Kk_zyo},
			},
		},
	},
	"t": {
		Basic: KanaGroup{
			Monographs: List["ta"],
			Digraphs: KanaRow{
				{"tya", "cha", Hg_tya, Kk_tya},
				{"tyu", "chu", Hg_tyu, Kk_tyu},
				{"tyo", "cho", Hg_tyo, Kk_tyo},
			},
		},
		Dakuten: KanaGroup{
			Monographs: List["da"],
			Digraphs: KanaRow{
				{"dya", "ja", Hg_dya, Kk_dya},
				{"dyu", "ju", Hg_dyu, Kk_dyu},
				{"dyo", "jo", Hg_dyo, Kk_dyo},
			},
		},
	},
	"n": {
		Basic: KanaGroup{
			Monographs: List["na"],
			Digraphs: KanaRow{
				{"nya", "", Hg_nya, Kk_nya},
				{"nyu", "", Hg_nyu, Kk_nyu},
				{"nyo", "", Hg_nyo, Kk_nyo},
			},
		},
	},
	"h": {
		Basic: KanaGroup{
			Monographs: List["ha"],
			Digraphs: KanaRow{
				{"hya", "", Hg_hya, Kk_hya},
				{"hyu", "", Hg_hyu, Kk_hyu},
				{"hyo", "", Hg_hyo, Kk_hyo},
			},
		},
		Dakuten: KanaGroup{
			Monographs: List["ba"],
			Digraphs: KanaRow{
				{"bya", "", Hg_bya, Kk_bya},
				{"byu", "", Hg_byu, Kk_byu},
				{"byo", "", Hg_byo, Kk_byo},
			},
		},
		Handakuten: KanaGroup{
			Monographs: List["pa"],
			Digraphs: KanaRow{
				{"pya", "", Hg_pya, Kk_pya},
				{"pyu", "", Hg_pyu, Kk_pyu},
				{"pyo", "", Hg_pyo, Kk_pyo},
			},
		},
	},
	"m": {
		Basic: KanaGroup{
			Monographs: List["ma"],
			Digraphs: KanaRow{
				{"mya", "", Hg_mya, Kk_mya},
				{"myu", "", Hg_myu, Kk_myu},
				{"myo", "", Hg_myo, Kk_myo},
			},
		},
	},
	"y": {
		Basic: KanaGroup{
			Monographs: List["ya"],
		},
	},
	"r": {
		Basic: KanaGroup{
			Monographs: List["ra"],
			Digraphs: KanaRow{
				{"rya", "", Hg_rya, Kk_rya},
				{"ryu", "", Hg_ryu, Kk_ryu},
				{"ryo", "", Hg_ryo, Kk_ryo},
			},
		},
	},
	"w": {
		Basic: KanaGroup{
			Monographs: KanaRow{
				{"wa", "", Hg_wa, Kk_wa},
				{"wo", "", Hg_wo, Kk_wo},
			},
		},
	},
	// Should we call this "N" instead of "nn"?
	"nn": {
		Basic: KanaGroup{
			Monographs: KanaRow{
				{"n", "nn", Hg_n, Kk_n},
			},
		},
	},
}

// GetHiragana returns the list of hiragana glyphs for the given row.
func GetHiragana(row KanaRow) []Hiragana {
	hiragana := []Hiragana{}
	for _, kana := range row {
		hiragana = append(hiragana, kana.Hiragana)
	}
	return hiragana
}

// GetKatakana returns the list of katakana glyphs for the given row.
func GetKatakana(row KanaRow) []Katakana {
	katakana := []Katakana{}
	for _, kana := range row {
		katakana = append(katakana, kana.Katakana)
	}
	return katakana
}

// GetRomaji returns the list of romaji strings for the given row.
func GetRomaji(row KanaRow) []Romaji {
	romaji := []Romaji{}
	for _, kana := range row {
		romaji = append(romaji, kana.Romaji)
	}
	return romaji
}
