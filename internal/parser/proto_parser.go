package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/lexer"
	"github.com/93mmm/proto-parser/internal/source"
)

type protoParser struct {
	*lexer.Lexer
}

func NewProtoParser(l *lexer.Lexer) *protoParser {
	return &protoParser{
		Lexer: l,
	}
}

func newTestProtoParser(in string) *protoParser {
	src := source.NewStringSource(in)
	bp := base.NewBaseParser(src)
	l := lexer.NewLexer(bp)
	return NewProtoParser(l)
}
