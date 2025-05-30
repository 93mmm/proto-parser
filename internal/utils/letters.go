package utils

import "unicode"

func IsValidInProto(symbol rune) bool {
	return unicode.IsLetter(symbol) ||
		unicode.IsDigit(symbol) ||
		symbol == '_' ||
		symbol == ' ' ||
		symbol == '\t' ||
		symbol == '\n'
}
