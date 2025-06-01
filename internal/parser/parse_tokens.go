package parser

import (
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

// syntax = "proto3";
func (p *ProtoParser) ParseSyntaxToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Syntax)

	if err := p.peekEquals(); err != nil {
		return nil, err
	}

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractQuotedString()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	if err := p.peekSemicolon(); err != nil {
		return nil, err
	}
	return s, nil
}

// package example;
func (p *ProtoParser) ParsePackageToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Package)

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	if err := p.peekSemicolon(); err != nil {
		return nil, err
	}
	return s, nil
}

// import "google/protobuf/timestamp.proto";
func (p *ProtoParser) ParseImportToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Import)

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractQuotedString()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	if err := p.peekSemicolon(); err != nil {
		return nil, err
	}
	return s, nil
}

// option go_package = "gitlab.ozon.ru/example/api/example;example";
func (p *ProtoParser) ParseOptionToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Option)

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	if err := p.peekEquals(); err != nil {
		return nil, err
	}

	p.skipWhiteSpaces()
	if _, err := p.extractQuotedString(); err != nil {
		return nil, err
	}

	if err := p.peekSemicolon(); err != nil {
		return nil, err
	}
	return s, nil
}

//	service Example {
//	  rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
//	}
func (p *ProtoParser) ParseServiceToken() ([]*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Service)

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	p.skipWhiteSpaces()

	if err := p.peekOpenBrace(); err != nil {
		return nil, err
	}

	p.skipWhiteSpaces()
	rpcs := make([]*symbols.Symbol, 0, 4)
	rpcs = append(rpcs, s)

	for !p.Test('}') && !p.EOF() {
		keyword, _ := p.extractKeyword()
		if keyword != "rpc" {
			// TODO: check error
		}
		r, _ := p.ParseRpcToken() // TODO: check error
		rpcs = append(rpcs, r)
		p.skipWhiteSpaces()
	}

	return rpcs, nil
}

// rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
func (p *ProtoParser) ParseRpcToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Rpc)

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	if _, err := p.extractNameBetweenParentheses(); err != nil {
		return nil, err
	}
	p.extractKeyword()
	if _, err := p.extractNameBetweenParentheses(); err != nil {
		return nil, err
	}

	p.skipUntilMatch(';')
	if err := p.peekSemicolon(); err != nil {
		return nil, err
	}
	return s, nil
}

// enum ExampleEnum {
//   ONE = 0;
//   TWO = 1;
//   THREE = 2;
// }
func (p *ProtoParser) ParseEnumToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Enum)

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	p.skipUntilMatch('{')
	p.skipUntilMatch('}')
	return s, nil
}

// message ExampleRPCResponse {}
func (p *ProtoParser) ParseMessageToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Message)

	if err := p.peekSemicolon(); err != nil {
		return nil, err
	}
	return s, nil
}
