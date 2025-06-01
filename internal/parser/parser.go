package parser

import (
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

type Parser interface {
	ParseDocument() ([]*symbols.Symbol, error)
}

type parser struct {
	*protoParser
}

type Source interface {
	Next() (rune, error)
}

func NewParser(pp *protoParser) Parser {
	return &parser{
		protoParser: pp,
	}
}

func (p *parser) ParseDocument() ([]*symbols.Symbol, error) {
	symbols := make([]*symbols.Symbol, 0, 10)

	for !p.EOF() {
		word, err := p.ExtractKeyword()
		if err != nil {
			return nil, err
		}

		switch word {
		case token.Syntax:
			p, err := p.ParseSyntaxToken()
			if err != nil {
				return nil, err
			}
			symbols = append(symbols, p)
		case token.Package:
			p, err := p.ParsePackageToken()
			if err != nil {
				return nil, err
			}
			symbols = append(symbols, p)
		case token.Import:
			p, err := p.ParseImportToken()
			if err != nil {
				return nil, err
			}
			symbols = append(symbols, p)
		case token.Option:
			p, err := p.ParseOptionToken()
			if err != nil {
				return nil, err
			}
			symbols = append(symbols, p)
		case token.Service:
			p, err := p.ParseServiceToken()
			if err != nil {
				return nil, err
			}
			symbols = append(symbols, p...)
		case token.Enum:
			p, err := p.ParseEnumToken()
			if err != nil {
				return nil, err
			}
			symbols = append(symbols, p)
		case token.Message:
			p, err := p.ParseMessageToken()
			if err != nil {
				return nil, err
			}
			symbols = append(symbols, p)
		}
		p.SkipWhiteSpaces()
	}
	return symbols, nil
}
