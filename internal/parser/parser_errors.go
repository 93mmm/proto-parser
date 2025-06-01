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

func NewParserError(line int, char int, msg string, args ...any) *ParserError {
	return &ParserError{
		message: fmt.Sprintf(msg, args...),
		lineNumber: line,
		charNumber: char,
	}
}

