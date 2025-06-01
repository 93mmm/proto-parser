package parser

import (
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
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
		word, err := p.ExtractKeyword()
		if err != nil {
			return nil, err
		}

		switch word {
		case token.Syntax:
			err := p.ParseSyntaxToken(syms)
			if err != nil {
				return nil, err
			}
		case token.Package:
			err := p.ParsePackageToken(syms)
			if err != nil {
				return nil, err
			}
		case token.Import:
			err := p.ParseImportToken(syms)
			if err != nil {
				return nil, err
			}
		case token.Option:
			err := p.ParseOptionToken(syms)
			if err != nil {
				return nil, err
			}
		case token.Service:
			err := p.ParseServiceToken(syms)
			if err != nil {
				return nil, err
			}
		case token.Enum:
			err := p.ParseEnumToken(syms)
			if err != nil {
				return nil, err
			}
		case token.Message:
			err := p.ParseMessageToken(syms)
			if err != nil {
				return nil, err
			}
		}
		p.SkipWhiteSpaces()
	}
	return syms.All(), nil
}
