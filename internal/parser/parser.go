package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/source"
)

type ProtoParser struct {
	base.BaseParser
}

func NewProtoParser(src source.Source) *ProtoParser {
	return &ProtoParser{
		BaseParser: *base.NewBaseParser(src),
	}
}

func (p *ProtoParser) extractWord() string {
	word := make([]rune, 0, 30)

	for !p.EOF() {
		if p.PeekSymbol() {
			word = append(word, p.Next())
		} else {
			break
		}
	}
	return string(word)
}

func (p *ProtoParser) skipWhiteSpaces() {
	for !p.EOF() {
		if p.PeekWhiteSpace() {
			p.Next()
		} else {
			break
		}
	}
}
