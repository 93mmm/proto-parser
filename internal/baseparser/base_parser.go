package baseparser

import (
	"github.com/93mmm/proto-parser/internal/source"
)

type BaseParser struct {
	src         source.Source
	currentChar rune
	eof         bool
}

func NewBaseParser(src source.Source) *BaseParser {
	p := &BaseParser{
		src: src,
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
	p.currentChar = next
	return r
}

func (p *BaseParser) EOF() bool {
	return p.eof
}

func (p *BaseParser) Test(expected rune) bool {
	return expected == p.currentChar
}
