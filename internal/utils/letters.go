package utils

import "unicode"

func IsValidInProto(symbol rune) bool {
	return unicode.IsLetter(symbol) || unicode.IsDigit(symbol) || symbol == '_'
}
