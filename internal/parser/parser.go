package parser

import (
	"github.com/93mmm/proto-parser/internal/errors"
	"github.com/93mmm/proto-parser/internal/symbols"
)

type Parser interface {
	ParseDocument() ([]*symbols.Symbol, error)
}

type parser struct {
	*tokenParser
}

type Source interface {
	Next() (rune, error)
}

func NewParser(pp *tokenParser) Parser {
	return &parser{
		tokenParser: pp,
	}
}

func (p *parser) ParseDocument() ([]*symbols.Symbol, error) {
	syms := symbols.NewCollector(10)

	for !p.EOF() {
		keyword, err := p.ExtractKeyword()
		if err != nil {
			return nil, err
		}
		b, ok := getBuilder(keyword)
		if !ok {
			return nil, errors.NewError(p.LineNumber(), p.CharNumber(), "Invalid keyword received")
		}
		err = b.Parse(p.tokenParser, syms)
		p.SkipWhiteSpaces()
	}

	return syms.All(), nil
}
