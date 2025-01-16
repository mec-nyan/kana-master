--[[ Helper nvim functions to ensure we don't write the kanas wrong.
--   And also so we don't have to type it all by hand!
--]]

-- to_hiragana yanks the silable under the cursor and, if it is a valid
-- romaji, it declares the corresponding hiragana symbol.
function to_hiragana()
	local word = vim.fn.expand('<cword>')
	local cmd = "ihg_<ESC>A Hiragana = '<C-k>" .. word
	if string.match(word, '^[aiueon]$') ~= nil then
		cmd = cmd .. '5'
	end
	cmd = cmd .. "'<ESC>"
	local cmd = vim.api.nvim_replace_termcodes(cmd, true, true, true)
	vim.api.nvim_feedkeys(cmd, 'n', false)
end

-- to_kana takes a romaji and a suffix (see :h digraphs) and tries to declare
-- the corresponding Hiragana or Katakana symbol.
-- suffix should be '5' (string, not number) for hiragana and '6' for katakana.
function to_kana(word, suffix)
	local prefix = "hg_"
	local _type = "Hiragana"
	if suffix == '6' then
		prefix = 'kk_'
		_type = "Katakana"
	end
	local cmd = "o".. prefix .. word .. "<ESC>A " .. _type .. " = '<C-k>" .. word
	if string.match(word, '^[aiueonAIUEON]$') ~= nil then
		cmd = cmd .. suffix
	end
	cmd = cmd .. "'<ESC>"
	local cmd = vim.api.nvim_replace_termcodes(cmd, true, true, true)
	vim.api.nvim_feedkeys(cmd, 'n', false)
end

-- to_kana_split takes a string of space separated romajis and passes them to "to_kana".
-- See "to_kana" for the suffix.
function to_kana_split(str, suffix)
	for i in string.gmatch(str, "%S+") do
		to_kana(i, suffix)
	end
end

-- to_kana_full takes a string of separated romaji introductors (representing the kana line) and adds a vowel for each corresponding column.
-- It then add the declaration for each kana.
function to_kana_full(str, suffix)
	local cmd = "o<ESC>"
	local cmd = vim.api.nvim_replace_termcodes(cmd, true, true, true)
	for i in string.gmatch(str, "%S+") do
		local whole = ""
		if i == 'y' then
			whole = 'ya yu yo'
		elseif i == 'Y' then
			whole = 'Ya Yu Yo'
		elseif i == 'w' then
			whole = 'wa wo n'
		elseif i == 'W' then
			whole = 'Wa Wo N'
		elseif i == 'a' then
			whole = 'a ' .. 'i ' .. 'u ' .. 'e ' .. 'o'
		elseif i == 'A' then
			whole = 'A ' .. 'I ' .. 'U ' .. 'E ' .. 'O'
		else
			whole = i .. 'a ' .. i .. 'i ' .. i .. 'u ' .. i .. 'e ' .. i .. 'o'
		end
		to_kana_split(whole, suffix)
		vim.api.nvim_feedkeys(cmd, 'n', false)
	end
end

-- print_hiragana declares all hiragana symbols (except compounds like "kya" etc).
function print_hiragana()
	local lines = 'a k s t n h m y r w g z d b p'
	to_kana_full(lines, '5')
end

-- print_hiragana declares all katakana symbols (except compounds like "kya" etc).
function print_katakana()
	local lines = 'A K S T N H M Y R W G Z D B P'
	to_kana_full(lines, '6')
end

