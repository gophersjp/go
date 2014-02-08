package doc

import (
	"strings"
	"unicode"
)

type japanese struct {
}

func (f japanese) isValidHead(r rune) bool {
	hiragana := unicode.Range16{0x3041, 0x3093, 1}
	katakana := unicode.Range16{0x30A1, 0x30F6, 1}
	chineseCharactors := unicode.Range16{0x4E00, 0x9FA0, 1}
	rt := unicode.RangeTable{R16: []unicode.Range16{hiragana, katakana, chineseCharactors}}
	if unicode.Is(&rt, r) {
		return true
	}
	return false
	
}

func (f japanese) isValidEnd(r rune) bool {
	hiragana := unicode.Range16{0x3041, 0x3093, 1}
	katakana := unicode.Range16{0x30A1, 0x30F6, 1}
	chineseCharactors := unicode.Range16{0x4E00, 0x9FA0, 1}
	rt := unicode.RangeTable{R16: []unicode.Range16{hiragana, katakana, chineseCharactors}}
	if unicode.Is(&rt, r) {
		return true
	}
	return false
	
}

func (f japanese) hasInvalidCharacters(line string) bool {
	if strings.IndexAny(line, "、。；：！？＋＊／＝（）「」｛｝＿＾＆〜％＃＠＜”＞￥") >= 0 {
		return true
	}
	return false
}
