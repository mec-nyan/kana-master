// Package kana associates hiragana/katakana with romaji.
package kana

type (
	Hiragana rune
	Katakana rune
	Digraph  string
	Romaji   string
)

const (
	Hg_a Hiragana = 'あ'
	Hg_i Hiragana = 'い'
	Hg_u Hiragana = 'う'
	Hg_e Hiragana = 'え'
	Hg_o Hiragana = 'お'

	Hg_ka Hiragana = 'か'
	Hg_ki Hiragana = 'き'
	Hg_ku Hiragana = 'く'
	Hg_ke Hiragana = 'け'
	Hg_ko Hiragana = 'こ'

	Hg_sa Hiragana = 'さ'
	Hg_si Hiragana = 'し'
	Hg_su Hiragana = 'す'
	Hg_se Hiragana = 'せ'
	Hg_so Hiragana = 'そ'

	Hg_ta Hiragana = 'た'
	Hg_ti Hiragana = 'ち'
	Hg_tu Hiragana = 'つ'
	Hg_te Hiragana = 'て'
	Hg_to Hiragana = 'と'

	Hg_na Hiragana = 'な'
	Hg_ni Hiragana = 'に'
	Hg_nu Hiragana = 'ぬ'
	Hg_ne Hiragana = 'ね'
	Hg_no Hiragana = 'の'

	Hg_ha Hiragana = 'は'
	Hg_hi Hiragana = 'ひ'
	Hg_hu Hiragana = 'ふ'
	Hg_he Hiragana = 'へ'
	Hg_ho Hiragana = 'ほ'

	Hg_ma Hiragana = 'ま'
	Hg_mi Hiragana = 'み'
	Hg_mu Hiragana = 'む'
	Hg_me Hiragana = 'め'
	Hg_mo Hiragana = 'も'

	Hg_ya Hiragana = 'や'
	Hg_yu Hiragana = 'ゆ'
	Hg_yo Hiragana = 'よ'

	Hg_ra Hiragana = 'ら'
	Hg_ri Hiragana = 'り'
	Hg_ru Hiragana = 'る'
	Hg_re Hiragana = 'れ'
	Hg_ro Hiragana = 'ろ'

	Hg_wa Hiragana = 'わ'
	Hg_wo Hiragana = 'を'
	Hg_n  Hiragana = 'ん'

	Hg_ga Hiragana = 'が'
	Hg_gi Hiragana = 'ぎ'
	Hg_gu Hiragana = 'ぐ'
	Hg_ge Hiragana = 'げ'
	Hg_go Hiragana = 'ご'

	Hg_za Hiragana = 'ざ'
	Hg_zi Hiragana = 'じ'
	Hg_zu Hiragana = 'ず'
	Hg_ze Hiragana = 'ぜ'
	Hg_zo Hiragana = 'ぞ'

	Hg_da Hiragana = 'だ'
	Hg_di Hiragana = 'ぢ'
	Hg_du Hiragana = 'づ'
	Hg_de Hiragana = 'で'
	Hg_do Hiragana = 'ど'

	Hg_ba Hiragana = 'ば'
	Hg_bi Hiragana = 'び'
	Hg_bu Hiragana = 'ぶ'
	Hg_be Hiragana = 'べ'
	Hg_bo Hiragana = 'ぼ'

	Hg_pa Hiragana = 'ぱ'
	Hg_pi Hiragana = 'ぴ'
	Hg_pu Hiragana = 'ぷ'
	Hg_pe Hiragana = 'ぺ'
	Hg_po Hiragana = 'ぽ'

	// Soma aliases for other common romaji forms.
	// Warning: Some conflicting forms are not being included.
	// i.e. "ji" and "zu"
	Hg_shi Hiragana = Hg_si
	Hg_chi Hiragana = Hg_ti
	Hg_tsu Hiragana = Hg_tu
	Hg_fu  Hiragana = Hg_hu
	Hg_ji  Hiragana = Hg_zi

	// Digraphs.
	// TODO: Should we use strings or []Hiragana?
	Hg_kya Digraph = "きゃ"
	Hg_kyu Digraph = "きゅ"
	Hg_kyo Digraph = "きょ"

	Hg_sha Digraph = "しゃ"
	Hg_shu Digraph = "しゅ"
	Hg_sho Digraph = "しょ"

	Hg_cha Digraph = "ちゃ"
	Hg_chu Digraph = "ちゅ"
	Hg_cho Digraph = "ちょ"

	Hg_nya Digraph = "にゃ"
	Hg_nyu Digraph = "にゅ"
	Hg_nyo Digraph = "にょ"

	Hg_mya Digraph = "みゃ"
	Hg_myu Digraph = "みゅ"
	Hg_myo Digraph = "みょ"

	Hg_rya Digraph = "りゃ"
	Hg_ryu Digraph = "りゅ"
	Hg_ryo Digraph = "りょ"

	Hg_gya Digraph = "ぎゃ"
	Hg_gyu Digraph = "ぎゅ"
	Hg_gyo Digraph = "ぎょ"

	Hg_sya Digraph = "じゃ"
	Hg_syu Digraph = "じゅ"
	Hg_syo Digraph = "じょ"

	Hg_tya Digraph = "ちゃ"
	Hg_tyu Digraph = "ちゅ"
	Hg_tyo Digraph = "ちょ"

	Hg_bya Digraph = "びゃ"
	Hg_byu Digraph = "びゅ"
	Hg_byo Digraph = "びょ"

	Hg_pya Digraph = "ぴゃ"
	Hg_pyu Digraph = "ぴゅ"
	Hg_pyo Digraph = "ぴょ"

	// End Hiragana.

	Kk_a Katakana = 'ア'
	Kk_i Katakana = 'イ'
	Kk_u Katakana = 'ウ'
	Kk_e Katakana = 'エ'
	Kk_o Katakana = 'オ'

	Kk_ka Katakana = 'カ'
	Kk_ki Katakana = 'キ'
	Kk_ku Katakana = 'ク'
	Kk_ke Katakana = 'ケ'
	Kk_ko Katakana = 'コ'

	Kk_sa Katakana = 'サ'
	Kk_si Katakana = 'シ'
	Kk_su Katakana = 'ス'
	Kk_se Katakana = 'セ'
	Kk_so Katakana = 'ソ'

	Kk_ta Katakana = 'タ'
	Kk_ti Katakana = 'チ'
	Kk_tu Katakana = 'ツ'
	Kk_te Katakana = 'テ'
	Kk_to Katakana = 'ト'

	Kk_na Katakana = 'ナ'
	Kk_ni Katakana = 'ニ'
	Kk_nu Katakana = 'ヌ'
	Kk_ne Katakana = 'ネ'
	Kk_no Katakana = 'ノ'

	Kk_ha Katakana = 'ハ'
	Kk_hi Katakana = 'ヒ'
	Kk_hu Katakana = 'フ'
	Kk_he Katakana = 'ヘ'
	Kk_ho Katakana = 'ホ'

	Kk_ma Katakana = 'マ'
	Kk_mi Katakana = 'ミ'
	Kk_mu Katakana = 'ム'
	Kk_me Katakana = 'メ'
	Kk_mo Katakana = 'モ'

	Kk_ya Katakana = 'ヤ'
	Kk_yu Katakana = 'ユ'
	Kk_yo Katakana = 'ヨ'

	Kk_ra Katakana = 'ラ'
	Kk_ri Katakana = 'リ'
	Kk_ru Katakana = 'ル'
	Kk_re Katakana = 'レ'
	Kk_ro Katakana = 'ロ'

	Kk_wa Katakana = 'ワ'
	Kk_wo Katakana = 'ヲ'
	Kk_n  Katakana = 'ン'

	Kk_ga Katakana = 'ガ'
	Kk_gi Katakana = 'ギ'
	Kk_gu Katakana = 'グ'
	Kk_ge Katakana = 'ゲ'
	Kk_go Katakana = 'ゴ'

	Kk_za Katakana = 'ザ'
	Kk_zi Katakana = 'ジ'
	Kk_zu Katakana = 'ズ'
	Kk_ze Katakana = 'ゼ'
	Kk_zo Katakana = 'ゾ'

	Kk_da Katakana = 'ダ'
	Kk_di Katakana = 'ヂ'
	Kk_du Katakana = 'ヅ'
	Kk_de Katakana = 'デ'
	Kk_do Katakana = 'ド'

	Kk_ba Katakana = 'バ'
	Kk_bi Katakana = 'ビ'
	Kk_bu Katakana = 'ブ'
	Kk_be Katakana = 'ベ'
	Kk_bo Katakana = 'ボ'

	Kk_pa Katakana = 'パ'
	Kk_pi Katakana = 'ピ'
	Kk_pu Katakana = 'プ'
	Kk_pe Katakana = 'ペ'
	Kk_po Katakana = 'ポ'

	// Soma aliases for other common romaji forms.
	// Warning: Some conflicting forms are not being included.
	// i.e. "ji" and "zu"
	Kk_shi Katakana = Kk_si
	Kk_chi Katakana = Kk_ti
	Kk_tsu Katakana = Kk_tu
	Kk_fu  Katakana = Kk_hu
	Kk_ji  Katakana = Kk_zi

	// Digraphs.
	// TODO: Should we use strings or []Katakana?
	Kk_kya Digraph = "キャ"
	Kk_kyu Digraph = "キュ"
	Kk_kyo Digraph = "キョ"

	Kk_sha Digraph = "シャ"
	Kk_shu Digraph = "シュ"
	Kk_sho Digraph = "ショ"

	Kk_cha Digraph = "チャ"
	Kk_chu Digraph = "シュ"
	Kk_cho Digraph = "チョ"

	Kk_nya Digraph = "ニャ"
	Kk_nyu Digraph = "ニュ"
	Kk_nyo Digraph = "ニョ"

	Kk_hya Digraph = "ヒャ"
	Kk_hyu Digraph = "ヒュ"
	Kk_hyo Digraph = "ヒョ"

	Kk_mya Digraph = "ミャ"
	Kk_myu Digraph = "みゅ"
	Kk_myo Digraph = "ミョ"

	Kk_rya Digraph = "リャ"
	Kk_ryu Digraph = "リュ"
	Kk_ryo Digraph = "リョ"

	Kk_gya Digraph = "ギャ"
	Kk_gyu Digraph = "ギュ"
	Kk_gyo Digraph = "ギョ"

	Kk_sya Digraph = "シャ"
	Kk_syu Digraph = "シュ"
	Kk_syo Digraph = "ショ"

	Kk_tya Digraph = "チャ"
	Kk_tyu Digraph = "チュ"
	Kk_tyo Digraph = "チョ"

	Kk_bya Digraph = "ビャ"
	Kk_byu Digraph = "ビュ"
	Kk_byo Digraph = "ビョ"

	Kk_pya Digraph = "ピャ"
	Kk_pyu Digraph = "ピュ"
	Kk_pyo Digraph = "ピョ"
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
