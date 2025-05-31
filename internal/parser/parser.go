package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

type ProtoParser struct {
	base.BaseParser
	tokens map[string]func() (*symbols.Symbol, error)
}

func NewProtoParser(src source.Source) *ProtoParser {
	p := &ProtoParser{
		BaseParser: *base.NewBaseParser(src),
	}

	p.tokens = map[string]func() (*symbols.Symbol, error) {
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
		word, _ := p.extractKeyword() // TODO: check error

		parsed, err := p.tokens[word]()
		if err != nil {
			panic(err) // TODO: not panic here!
		}
		symbols = append(symbols, parsed)
	}
	return symbols
}
