// Package kana associates hiragana/katakana with romaji.
package kana

type (
	Hiragana rune
	Katakana rune
	Romaji   string
)

const (
	hg_a Hiragana = 'あ'
	hg_i Hiragana = 'い'
	hg_u Hiragana = 'う'
	hg_e Hiragana = 'え'
	hg_o Hiragana = 'お'

	hg_ka Hiragana = 'か'
	hg_ki Hiragana = 'き'
	hg_ku Hiragana = 'く'
	hg_ke Hiragana = 'け'
	hg_ko Hiragana = 'こ'

	hg_sa Hiragana = 'さ'
	hg_si Hiragana = 'し'
	hg_su Hiragana = 'す'
	hg_se Hiragana = 'せ'
	hg_so Hiragana = 'そ'

	hg_ta Hiragana = 'た'
	hg_ti Hiragana = 'ち'
	hg_tu Hiragana = 'つ'
	hg_te Hiragana = 'て'
	hg_to Hiragana = 'と'

	hg_na Hiragana = 'な'
	hg_ni Hiragana = 'に'
	hg_nu Hiragana = 'ぬ'
	hg_ne Hiragana = 'ね'
	hg_no Hiragana = 'の'

	hg_ha Hiragana = 'は'
	hg_hi Hiragana = 'ひ'
	hg_hu Hiragana = 'ふ'
	hg_he Hiragana = 'へ'
	hg_ho Hiragana = 'ほ'

	hg_ma Hiragana = 'ま'
	hg_mi Hiragana = 'み'
	hg_mu Hiragana = 'む'
	hg_me Hiragana = 'め'
	hg_mo Hiragana = 'も'

	hg_ya Hiragana = 'や'
	hg_yu Hiragana = 'ゆ'
	hg_yo Hiragana = 'よ'

	hg_ra Hiragana = 'ら'
	hg_ri Hiragana = 'り'
	hg_ru Hiragana = 'る'
	hg_re Hiragana = 'れ'
	hg_ro Hiragana = 'ろ'

	hg_wa Hiragana = 'わ'
	hg_wo Hiragana = 'を'
	hg_n Hiragana = 'ん'

	hg_ga Hiragana = 'が'
	hg_gi Hiragana = 'ぎ'
	hg_gu Hiragana = 'ぐ'
	hg_ge Hiragana = 'げ'
	hg_go Hiragana = 'ご'

	hg_za Hiragana = 'ざ'
	hg_zi Hiragana = 'じ'
	hg_zu Hiragana = 'ず'
	hg_ze Hiragana = 'ぜ'
	hg_zo Hiragana = 'ぞ'

	hg_da Hiragana = 'だ'
	hg_di Hiragana = 'ぢ'
	hg_du Hiragana = 'づ'
	hg_de Hiragana = 'で'
	hg_do Hiragana = 'ど'

	hg_ba Hiragana = 'ば'
	hg_bi Hiragana = 'び'
	hg_bu Hiragana = 'ぶ'
	hg_be Hiragana = 'べ'
	hg_bo Hiragana = 'ぼ'

	hg_pa Hiragana = 'ぱ'
	hg_pi Hiragana = 'ぴ'
	hg_pu Hiragana = 'ぷ'
	hg_pe Hiragana = 'ぺ'
	hg_po Hiragana = 'ぽ'

	// Soma aliases for other common romaji forms.
	// Warning: Some conflicting forms are not being included.
	// i.e. "ji" and "zu"
	hg_shi Hiragana = hg_si
	hg_chi Hiragana = hg_ti
	hg_tsu Hiragana = hg_tu
	hg_fo Hiragana = hg_hu
	hg_ji Hiragana = hg_zi

	// End Hiragana.

	kk_a Katakana = 'ア'
	kk_i Katakana = 'イ'
	kk_u Katakana = 'ウ'
	kk_e Katakana = 'エ'
	kk_o Katakana = 'オ'

	kk_ka Katakana = 'カ'
	kk_ki Katakana = 'キ'
	kk_ku Katakana = 'ク'
	kk_ke Katakana = 'ケ'
	kk_ko Katakana = 'コ'

	kk_sa Katakana = 'サ'
	kk_si Katakana = 'シ'
	kk_su Katakana = 'ス'
	kk_se Katakana = 'セ'
	kk_so Katakana = 'ソ'

	kk_ta Katakana = 'タ'
	kk_ti Katakana = 'チ'
	kk_tu Katakana = 'ツ'
	kk_te Katakana = 'テ'
	kk_to Katakana = 'ト'

	kk_na Katakana = 'ナ'
	kk_ni Katakana = 'ニ'
	kk_nu Katakana = 'ヌ'
	kk_ne Katakana = 'ネ'
	kk_no Katakana = 'ノ'

	kk_ha Katakana = 'ハ'
	kk_hi Katakana = 'ヒ'
	kk_hu Katakana = 'フ'
	kk_he Katakana = 'ヘ'
	kk_ho Katakana = 'ホ'

	kk_ma Katakana = 'マ'
	kk_mi Katakana = 'ミ'
	kk_mu Katakana = 'ム'
	kk_me Katakana = 'メ'
	kk_mo Katakana = 'モ'

	kk_ya Katakana = 'ヤ'
	kk_yu Katakana = 'ユ'
	kk_yo Katakana = 'ヨ'

	kk_ra Katakana = 'ラ'
	kk_ri Katakana = 'リ'
	kk_ru Katakana = 'ル'
	kk_re Katakana = 'レ'
	kk_ro Katakana = 'ロ'

	kk_wa Katakana = 'ワ'
	kk_wo Katakana = 'ヲ'
	kk_n Katakana = 'ン'

	kk_ga Katakana = 'ガ'
	kk_gi Katakana = 'ギ'
	kk_gu Katakana = 'グ'
	kk_ge Katakana = 'ゲ'
	kk_go Katakana = 'ゴ'

	kk_za Katakana = 'ザ'
	kk_zi Katakana = 'ジ'
	kk_zu Katakana = 'ズ'
	kk_ze Katakana = 'ゼ'
	kk_zo Katakana = 'ゾ'

	kk_da Katakana = 'ダ'
	kk_di Katakana = 'ヂ'
	kk_du Katakana = 'ヅ'
	kk_de Katakana = 'デ'
	kk_do Katakana = 'ド'

	kk_ba Katakana = 'バ'
	kk_bi Katakana = 'ビ'
	kk_bu Katakana = 'ブ'
	kk_be Katakana = 'ベ'
	kk_bo Katakana = 'ボ'

	kk_pa Katakana = 'パ'
	kk_pi Katakana = 'ピ'
	kk_pu Katakana = 'プ'
	kk_pe Katakana = 'ペ'
	kk_po Katakana = 'ポ'

	// Soma aliases for other common romaji forms.
	// Warning: Some conflicting forms are not being included.
	// i.e. "ji" and "zu"
	kk_shi Katakana = kk_si
	kk_chi Katakana = kk_ti
	kk_tsu Katakana = kk_tu
	kk_fo Katakana = kk_hu
	kk_ji Katakana = kk_zi
)

type Kana struct {
	Hiragana
	Katakana
}

var (
	hiragana map[Romaji]Hiragana
	katakana map[Romaji]Katakana
	kana     map[Romaji]Kana
)

