package kana

// The name of the constants use the "typing" form, since
// romanisation is not standard and sometimes it uses the same
// Romaji for different kanas.
const (
	// 五十音（ごじゅうおん）Monographs.
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

	// 濁点（だくてん）dakuten (diacritics).
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

	// 半濁点　（はんだくてん）handakuten (diacritics)
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
	// 五十音（ごじゅうおん）
	Kk_kya Digraph = "キャ"
	Kk_kyu Digraph = "キュ"
	Kk_kyo Digraph = "キョ"

	// Usually romanised as "sha, shu, sho"
	Kk_sya Digraph = "シャ"
	Kk_syu Digraph = "シュ"
	Kk_syo Digraph = "ショ"

	// Usually romanised as "cha, chu, cho"
	Kk_tya Digraph = "チャ"
	Kk_tyu Digraph = "チュ"
	Kk_tyo Digraph = "チョ"

	Kk_nya Digraph = "ニャ"
	Kk_nyu Digraph = "ニュ"
	Kk_nyo Digraph = "ニョ"

	Kk_hya Digraph = "ヒャ"
	Kk_hyu Digraph = "ヒュ"
	Kk_hyo Digraph = "ヒョ"

	Kk_mya Digraph = "ミャ"
	Kk_myu Digraph = "ミュ"
	Kk_myo Digraph = "ミョ"

	Kk_rya Digraph = "リャ"
	Kk_ryu Digraph = "リュ"
	Kk_ryo Digraph = "リョ"

	// 濁点（だくてん）dakuten (diacritics).
	Kk_gya Digraph = "ギャ"
	Kk_gyu Digraph = "ギュ"
	Kk_gyo Digraph = "ギョ"

	// Usually romanised as "ja, ju, jo"
	Kk_zya Digraph = "ジャ"
	Kk_zyu Digraph = "シュ"
	Kk_zyo Digraph = "ショ"

	// Usually also romanised as "ja, ju, jo"
	Kk_dya Digraph = "ヂャ"
	Kk_dyu Digraph = "ヂュ"
	Kk_dyo Digraph = "ヂョ"

	Kk_bya Digraph = "ビャ"
	Kk_byu Digraph = "ビュ"
	Kk_byo Digraph = "ビョ"

	// 半濁点　（はんだくてん）handakuten (diacritics)
	Kk_pya Digraph = "ピャ"
	Kk_pyu Digraph = "ピュ"
	Kk_pyo Digraph = "ピョ"
)
