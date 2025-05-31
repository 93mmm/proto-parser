package parser

import (
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

// syntax = "proto3";
func (p *ProtoParser) ParseSyntaxToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Syntax)

	p.skipWhiteSpaces()
	if !p.Peek('=') {
		return nil, NewParserError("Expected = found nothing", p.LineNumber(), p.CharNumber())
	}

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractQuotedString()
	if err != nil {
		return nil, err
	}

	s.SetEndChar(p.CharNumber())
	s.SetName(name)

	if !p.Peek(';') {
		return nil, NewParserError("Expected ; found nothing", p.LineNumber(), p.CharNumber())
	}

	return s, nil
}

func (p *ProtoParser) ParsePackageToken() (*symbols.Symbol, error) {
	return nil, nil
}

func (p *ProtoParser) ParseImportToken() (*symbols.Symbol, error) {
	return nil, nil
}

func (p *ProtoParser) ParseOptionToken() (*symbols.Symbol, error) {
	return nil, nil
}

func (p *ProtoParser) ParseServiceToken() (*symbols.Symbol, error) {
	return nil, nil
}

func (p *ProtoParser) ParseRpcToken() (*symbols.Symbol, error) {
	return nil, nil
}

func (p *ProtoParser) ParseEnumToken() (*symbols.Symbol, error) {
	return nil, nil
}

func (p *ProtoParser) ParseMessageToken() (*symbols.Symbol, error) {
	return nil, nil
}
