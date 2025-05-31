package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

type ProtoParser struct {
	base.BaseParser
	tokens map[string]func() *symbols.Symbol
}

func NewProtoParser(src source.Source) *ProtoParser {
	p := &ProtoParser{
		BaseParser: *base.NewBaseParser(src),
	}

	p.tokens = map[string]func() *symbols.Symbol{
		token.Syntax:  p.ParseSyntaxToken,
		token.Package: p.ParsePackageToken,
		token.Import:  p.ParseImportToken,
		token.Option:  p.ParseOptionToken,
		token.Service: p.ParseServiceToken,
		token.Rpc:     p.ParseRpcToken,
		token.Enum:    p.ParseEnumToken,
		token.Message: p.ParseMessageToken,
	}

	return p
}

func (p *ProtoParser) ParseDocument() []*symbols.Symbol {
	symbols := make([]*symbols.Symbol, 0, 10)

	for !p.EOF() {
		p.skipWhiteSpaces()
		word, _ := p.extractKeyword() // TODO: check error
		symbols = append(symbols, p.tokens[word]()) // TODO: panic if nil func ofc
	}
	return symbols
}

func (p *ProtoParser) ParseSyntaxToken() *symbols.Symbol {
	return nil
}

func (p *ProtoParser) ParsePackageToken() *symbols.Symbol {
	return nil
}

func (p *ProtoParser) ParseImportToken() *symbols.Symbol {
	return nil
}

func (p *ProtoParser) ParseOptionToken() *symbols.Symbol {
	return nil
}

func (p *ProtoParser) ParseServiceToken() *symbols.Symbol {
	return nil
}

func (p *ProtoParser) ParseRpcToken() *symbols.Symbol {
	return nil
}

func (p *ProtoParser) ParseEnumToken() *symbols.Symbol {
	return nil
}

func (p *ProtoParser) ParseMessageToken() *symbols.Symbol {
	return nil
}
