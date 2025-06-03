package base

import (
	"unicode"

	"github.com/93mmm/proto-parser/internal/parser/constants"
)

type Source interface {
	Next() (rune, error)
}

type BaseParser struct {
	src         Source
	currentChar rune
	eof         bool
	lineNumber  int
	charNumber  int
}

func NewBaseParser(src Source) *BaseParser {
	p := &BaseParser{
		src:        src,
		lineNumber: 1,
		charNumber: 0,
	}
	p.Next()
	return p
}

func (p *BaseParser) Peek(expected rune) bool {
	if p.Test(expected) {
		p.Next()
		return true
	}
	return false
}

func (p *BaseParser) Next() rune {
	r := p.currentChar
	next, err := p.src.Next()
	if err != nil {
		p.eof = true
		p.currentChar = constants.EOF
	}

	p.incrementCharOrLineNumber(next)
	p.currentChar = next
	return r
}

func (p *BaseParser) EOF() bool {
	return p.eof
}

func (p *BaseParser) incrementCharOrLineNumber(c rune) {
	if p.EOF() {
		return
	}
	if c == constants.NextLine {
		p.lineNumber++
		p.charNumber = 0
		return
	}
	p.charNumber++
}

func (p *BaseParser) CurrentChar() rune { return p.currentChar }
func (p *BaseParser) LineNumber() int   { return p.lineNumber }
func (p *BaseParser) CharNumber() int   { return p.charNumber }

func (p *BaseParser) Test(expected rune) bool {
	return expected == p.currentChar
}

func (p *BaseParser) TestKeyword() bool {
	return unicode.IsLetter(p.currentChar) &&
		unicode.IsLower(p.currentChar)
}

func (p *BaseParser) TestName() bool {
	return unicode.IsLetter(p.currentChar) ||
		unicode.IsDigit(p.currentChar) ||
		p.currentChar == constants.Underscore
}

func (p *BaseParser) TestWhiteSpace() bool {
	return p.currentChar == constants.Space ||
		p.currentChar == constants.Tab ||
		p.currentChar == constants.NextLine
}
