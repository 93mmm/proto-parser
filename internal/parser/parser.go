package parser

import (
	"github.com/93mmm/proto-parser/internal/errors"
	"github.com/93mmm/proto-parser/internal/parser/builder"
	"github.com/93mmm/proto-parser/internal/symbols"
)

type Parser interface {
	ParseDocument() ([]*symbols.Symbol, error)
}

type parser struct {
	*builder.TokenParser
}

type Source interface {
	Next() (rune, error)
}

func NewParser(pp *builder.TokenParser) Parser {
	return &parser{
		TokenParser: pp,
	}
}

func (p *parser) ParseDocument() ([]*symbols.Symbol, error) {
	syms := symbols.NewCollector(10)

	for !p.EOF() {
		keyword, err := p.ExtractKeyword()
		if err != nil {
			return nil, err
		}
		b, ok := builder.GetBuilder(keyword)
		if !ok {
			return nil, errors.NewError(p.LineNumber(), p.CharNumber(), "Invalid keyword received")
		}
		err = b.Parse(p.TokenParser, syms)
		p.SkipWhiteSpaces()
	}

	return syms.All(), nil
}
