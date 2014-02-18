package doc

type language interface {
	isValidHeader(line string) bool
}

var languages = []language{ja{}}

func isValidNonAlphabeticalScriptHeader(line string) bool {
	for _, v := range languages {
		if v.isValidHeader(line) {
			return true
		}
		return false
	}
	return false
}
