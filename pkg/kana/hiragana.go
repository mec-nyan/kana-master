package kana

// The name of the constants use the "typing" form, since
// romanisation is not standard and sometimes it uses the same
// Romaji for different kanas.
const (
	// 五十音（ごじゅうおん）Monographs.
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

	// 濁点（だくてん）dakuten (diacritics).
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

	// 半濁点　（はんだくてん）handakuten (diacritics)
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
	// 五十音（ごじゅうおん）
	Hg_kya Digraph = "きゃ"
	Hg_kyu Digraph = "きゅ"
	Hg_kyo Digraph = "きょ"

	// Usually romanised as "sha, shu, sho"
	Hg_sya Digraph = "しゃ"
	Hg_syu Digraph = "しゅ"
	Hg_syo Digraph = "しょ"

	// Usually romanised as "cha, chu, cho"
	Hg_tya Digraph = "ちゃ"
	Hg_tyu Digraph = "ちゅ"
	Hg_tyo Digraph = "ちょ"

	Hg_nya Digraph = "にゃ"
	Hg_nyu Digraph = "にゅ"
	Hg_nyo Digraph = "にょ"

	Hg_hya Digraph = "ひゃ"
	Hg_hyu Digraph = "ひゅ"
	Hg_hyo Digraph = "ひょ"

	Hg_mya Digraph = "みゃ"
	Hg_myu Digraph = "みゅ"
	Hg_myo Digraph = "みょ"

	Hg_rya Digraph = "りゃ"
	Hg_ryu Digraph = "りゅ"
	Hg_ryo Digraph = "りょ"

	// 濁点（だくてん）dakuten (diacritics).
	Hg_gya Digraph = "ぎゃ"
	Hg_gyu Digraph = "ぎゅ"
	Hg_gyo Digraph = "ぎょ"

	// Usually romanised as "ja, ju, jo"
	Hg_zya Digraph = "じゃ"
	Hg_zyu Digraph = "じゅ"
	Hg_zyo Digraph = "じょ"

	// Usually also romanised as "ja, ju, jo"
	Hg_dya Digraph = "ちゃ"
	Hg_dyu Digraph = "ちゅ"
	Hg_dyo Digraph = "ちょ"

	Hg_bya Digraph = "びゃ"
	Hg_byu Digraph = "びゅ"
	Hg_byo Digraph = "びょ"

	// 半濁点　（はんだくてん）handakuten (diacritics)
	Hg_pya Digraph = "ぴゃ"
	Hg_pyu Digraph = "ぴゅ"
	Hg_pyo Digraph = "ぴょ"

	// End Hiragana.
)
