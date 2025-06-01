package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

type ProtoParser struct {
	base.BaseParser
}

func NewProtoParser(src source.Source) *ProtoParser {
	p := &ProtoParser{
		BaseParser: *base.NewBaseParser(src),
	}

	return p
}

func (p *ProtoParser) ParseDocument() []*symbols.Symbol {
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
			p, err := p.ParseEnumToken()
			if err != nil {
				panic(err)
			}
			symbols = append(symbols, p)
		}
	}
	return symbols
}
