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

// Peek next if expected matches current
// Remain current if expected not matches current
func (p *BaseParser) Peek(expected rune) bool {
	if p.Test(expected) {
		p.Next()
		return true
	}
	return false
}

// Peek next
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

func (p *BaseParser) PeekSymbol() bool { return utils.IsSymbol(p.currentChar) }
func (p *BaseParser) PeekWhiteSpace() bool { return utils.IsWhiteSpace(p.currentChar) }
