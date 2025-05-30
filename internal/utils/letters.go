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

func IsSymbol(symbol rune) bool {
	return unicode.IsLetter(symbol) ||
		unicode.IsDigit(symbol) ||
		symbol == '_' ||
		symbol == '"'
}

func IsWhiteSpace(symbol rune) bool {
	return symbol == ' ' ||
		symbol == '\t' ||
		symbol == '\n'
}
