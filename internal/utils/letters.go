package utils

import "unicode"

func IsKeyword(symbol rune) bool {
	return unicode.IsLetter(symbol) &&
		unicode.IsLower(symbol)
}

func IsName(symbol rune) bool {
	return unicode.IsLetter(symbol) ||
		unicode.IsDigit(symbol) ||
		symbol == '_'

}

func IsQuote(symbol rune) bool {
	return symbol == '"'
}

func IsValidInProto(symbol rune) bool {
	return unicode.IsLetter(symbol) ||
		unicode.IsDigit(symbol) ||
		symbol == '_' ||
		symbol == ' ' ||
		symbol == '\t' ||
		symbol == '\n'
}

func IsWhiteSpace(symbol rune) bool {
	return symbol == ' ' ||
		symbol == '\t' ||
		symbol == '\n'
}
