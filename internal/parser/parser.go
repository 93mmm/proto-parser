package parser

import (
	"github.com/93mmm/proto-parser/internal/errors"
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

type TokenBuilder interface {
	Parse(*tokenParser, symbols.Collector) error
}

func NewParser(pp *tokenParser) Parser {
	return &parser{
		tokenParser: pp,
	}
}

func (p *parser) ParseDocument() ([]*symbols.Symbol, error) {
	syms := symbols.NewCollector(10)
	tokenBuilders := map[string]TokenBuilder{
		token.Syntax: SyntaxToken{},
		token.Package: PackageToken{},
		token.Import: ImportToken{},
		token.Option: OptionToken{},
		token.Service: ServiceToken{},
		token.Enum: EnumToken{},
		token.Message: MessageToken{},
	}

	for !p.EOF() {
		word, err := p.ExtractKeyword()
		if err != nil {
			return nil, err
		}
		b, ok := tokenBuilders[word]
		if !ok {
			return nil, errors.NewError(p.LineNumber(), p.CharNumber(), "Invalid keyword received")
		}
		err = b.Parse(p.tokenParser, syms)
		p.SkipWhiteSpaces()
	}

	return syms.All(), nil
}
