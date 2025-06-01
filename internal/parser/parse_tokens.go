package parser

import (
	"github.com/93mmm/proto-parser/internal/errors"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

// syntax = "proto3";
func (p *protoParser) ParseSyntaxToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Syntax)

	if err := p.PeekSymbol('='); err != nil {
		return nil, err
	}

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractQuotedString()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

// package example;
func (p *protoParser) ParsePackageToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Package)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

// import "google/protobuf/timestamp.proto";
func (p *protoParser) ParseImportToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Import)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractQuotedString()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

// option go_package = "gitlab.ozon.ru/example/api/example;example";
func (p *protoParser) ParseOptionToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Option)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol('='); err != nil {
		return nil, err
	}

	p.SkipWhiteSpaces()
	if _, err := p.ExtractQuotedString(); err != nil {
		return nil, err
	}

	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

//	service Example {
//	  rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
//	}
func (p *protoParser) ParseServiceToken() ([]*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Service)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	p.SkipWhiteSpaces()

	if err := p.PeekSymbol('{'); err != nil {
		return nil, err
	}

	p.SkipWhiteSpaces()
	rpcs := make([]*symbols.Symbol, 0, 4)
	rpcs = append(rpcs, s)

	for !p.Test('}') && !p.EOF() {
		keyword, _ := p.ExtractKeyword()
		if keyword != "rpc" {
			return nil, errors.NewParserError(p.LineNumber(), p.CharNumber(), "Unexpected keyword %s found inside %s", keyword, s)
		}
		r, err := p.ParseRpcToken()
		if err != nil {
			return nil, err
		}
		rpcs = append(rpcs, r)
		p.SkipWhiteSpaces()
	}
	p.Next()

	return rpcs, nil
}

// rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
func (p *protoParser) ParseRpcToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Rpc)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if _, err := p.ExtractNameBetweenParentheses(); err != nil {
		return nil, err
	}
	p.ExtractKeyword()
	if _, err := p.ExtractNameBetweenParentheses(); err != nil {
		return nil, err
	}

	p.SkipUntilMatch(';')
	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

//	enum ExampleEnum {
//	  ONE = 0;
//	  TWO = 1;
//	  THREE = 2;
//	}
func (p *protoParser) ParseEnumToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Enum)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	p.SkipCurlyBraces()
	return s, nil
}

// message ExampleRPCResponse {}
func (p *protoParser) ParseMessageToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Message)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	p.SkipCurlyBraces()
	return s, nil
}
