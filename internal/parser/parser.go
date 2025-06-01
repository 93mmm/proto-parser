package parser

import (
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

type Parser interface {
	ParseDocument() []*symbols.Symbol
}

type parser struct {
	protoParser
}

var _ Parser = (*parser)(nil)

func NewParser(src source.Source) Parser {
	return &parser{
		protoParser: *newProtoParser(src),
	}
}

func (p *parser) ParseDocument() []*symbols.Symbol {
	symbols := make([]*symbols.Symbol, 0, 10)

	for !p.EOF() {
		p.skipWhiteSpaces()
		word, _ := p.extractKeyword() // TODO: check error

		switch word {
		case token.Syntax:
			p, err := p.ParseSyntaxToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p)
		case token.Package:
			p, err := p.ParsePackageToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p)
		case token.Import:
			p, err := p.ParseImportToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p)
		case token.Option:
			p, err := p.ParseOptionToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p)
		case token.Service:
			p, err := p.ParseServiceToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p...)
		case token.Enum:
			p, err := p.ParseEnumToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p)
		case token.Message:
			p, err := p.ParseMessageToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p)
		}
	}
	return symbols
}
