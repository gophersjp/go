package doc

type language interface {
	isValidHead(r rune) bool
	isValidEnd(r rune) bool
	hasInvalidCharacters(line string) bool
}

var languages = []language{japanese{}}

func isValidNonLatinHeadCharacter(r rune) bool {
	for _, v := range languages {
		if v.isValidHead(r) {
			return true
		}
	}
	return false
}

func isValidNonLatinEndCharacter(r rune) bool {
	for _, v := range languages {
		if v.isValidEnd(r) {
			return true
		}
	}
	return false
}

func hasInvalidNonLatinCharacters(line string) bool {
	for _, v := range languages {
		if v.hasInvalidCharacters(line) {
			return true
		}
	}
	return false
}
