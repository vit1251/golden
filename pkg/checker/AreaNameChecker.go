package checker

var printable []rune = []rune{
	'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y',
	'Z', '.', '-', '~', '_'
}

func IsPrintable(source rune) bool {
	for _, ch := range printable {
		if source == ch {
			return true
		}
	}
	return false
}

func AreaName_Check(source string) bool {
	for _, ch := range source {
		if !IsPrintable(ch) {
			return false
		}
	}
	return true
}
