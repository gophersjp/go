package doc

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type ja struct {
}

func (f ja) isValidHeader(line string) bool {
	// a heading must start with a letter of Hiragana, Katakana, or Chinese characters
	r, _ := utf8.DecodeRuneInString(line)
	hiragana := unicode.Range16{0x3041, 0x3093, 1}
	katakana := unicode.Range16{0x30A1, 0x30F6, 1}
	chineseCharactors := unicode.Range16{0x4E00, 0x9FA0, 1}
	rt := unicode.RangeTable{R16: []unicode.Range16{hiragana, katakana, chineseCharactors}}
	if !unicode.Is(&rt, r) {
		return false
	}

	// it must end in a letter of Hiragana, Katakana, or Chinese characters
	r, _ = utf8.DecodeLastRuneInString(line)
	if !unicode.Is(&rt, r) {
		return false
	}

	// exclude lines with illegal characters
	if strings.IndexAny(line, "、。；：！？＋＊／＝（）「」｛｝＿＾＆〜％＃＠＜”＞￥") >= 0 {
		return false
	}

	return true
}
