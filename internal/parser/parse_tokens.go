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

	p.skipWhiteSpaces()
	name, err := p.extractQuotedString()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	p.skipWhiteSpaces()
	if !p.Peek(';') {
		return nil, NewParserError("Expected ; found nothing", p.LineNumber(), p.CharNumber())
	}

	return s, nil
}

// package example;
func (p *ProtoParser) ParsePackageToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Package)
	
	p.skipWhiteSpaces()
	name, err := p.extractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	p.skipWhiteSpaces()
	if !p.Peek(';') {
		return nil, NewParserError("Expected ; found nothing", p.LineNumber(), p.CharNumber())
	}

	return s, nil
}

func (p *ProtoParser) ParseImportToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Import)
	return s, nil
}

func (p *ProtoParser) ParseOptionToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Option)
	return s, nil
}

func (p *ProtoParser) ParseServiceToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Service)
	return s, nil
}

func (p *ProtoParser) ParseRpcToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Rpc)
	return s, nil
}

func (p *ProtoParser) ParseEnumToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Enum)
	return s, nil
}

func (p *ProtoParser) ParseMessageToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Message)
	return s, nil
}
