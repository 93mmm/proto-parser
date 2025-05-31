package baseparser

import (
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/utils"
)

type BaseParser struct {
	src         source.Source
	currentChar rune
	eof         bool
	lineNumber  int
	charNumber  int
}

func NewBaseParser(src source.Source) *BaseParser {
	p := &BaseParser{
		src:        src,
		lineNumber: 1,
		charNumber: 0,
	}
	p.Next()
	return p
}

// If expected matches current, return true and go forward
// Else return false and do nothing
func (p *BaseParser) Peek(expected rune) bool {
	if p.Test(expected) {
		p.Next()
		return true
	}
	return false
}

// Return current and go forward
func (p *BaseParser) Next() rune {
	r := p.currentChar
	next, err := p.src.Next()
	if err != nil {
		p.eof = true
		p.currentChar = source.EOF
	}

	p.incrementCharOrLineNumber(next)
	p.currentChar = next
	return r
}

func (p *BaseParser) EOF() bool {
	return p.eof
}

func (p *BaseParser) Test(expected rune) bool {
	return expected == p.currentChar
}

func (p *BaseParser) incrementCharOrLineNumber(c rune) {
	if p.EOF() {
		return
	}
	if c == '\n' {
		p.lineNumber++
		p.charNumber = 0
		return
	}
	p.charNumber++
}

func (p *BaseParser) LineNumber() int { return p.lineNumber }
func (p *BaseParser) CharNumber() int { return p.charNumber }

func (p *BaseParser) CurrentChar() rune { return p.currentChar }

func (p *BaseParser) TestQuote() bool      { return utils.IsQuote(p.currentChar) }
func (p *BaseParser) TestKeyword() bool    { return utils.IsKeyword(p.currentChar) }
func (p *BaseParser) TestName() bool       { return utils.IsName(p.currentChar) }
func (p *BaseParser) TestWhiteSpace() bool { return utils.IsWhiteSpace(p.currentChar) }
