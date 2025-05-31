package parser

import "fmt"

type ParserError struct {
	message    string
	lineNumber int
	charNumber int
}

func (e *ParserError) Error() string {
	return fmt.Sprintf(
		"ParserError with message: %v\nAt %v:%v",
		e.message,
		e.lineNumber,
		e.charNumber,
	)
}

func NewParserError(msg string, line int, char int) *ParserError {
	return &ParserError{
		message: msg,
		lineNumber: line,
		charNumber: char,
	}
}

