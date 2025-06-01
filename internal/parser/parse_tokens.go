package parser

import (
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

// syntax = "proto3";
func (p *protoParser) ParseSyntaxToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Syntax)

	if err := p.peekSymbol('='); err != nil {
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

	if err := p.peekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

// package example;
func (p *protoParser) ParsePackageToken() (*symbols.Symbol, error) {
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

	if err := p.peekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

// import "google/protobuf/timestamp.proto";
func (p *protoParser) ParseImportToken() (*symbols.Symbol, error) {
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

	if err := p.peekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

// option go_package = "gitlab.ozon.ru/example/api/example;example";
func (p *protoParser) ParseOptionToken() (*symbols.Symbol, error) {
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

	if err := p.peekSymbol('='); err != nil {
		return nil, err
	}

	p.skipWhiteSpaces()
	if _, err := p.extractQuotedString(); err != nil {
		return nil, err
	}

	if err := p.peekSymbol(';'); err != nil {
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

	if err := p.peekSymbol('{'); err != nil {
		return nil, err
	}

	p.skipWhiteSpaces()
	rpcs := make([]*symbols.Symbol, 0, 4)
	rpcs = append(rpcs, s)

	for !p.Test('}') && !p.EOF() {
		keyword, _ := p.extractKeyword()
		if keyword != "rpc" {
			return nil, NewParserError("Unexpected keyword " + keyword + "found inside " + s.String(), p.LineNumber(), p.CharNumber())
		}
		r, err := p.ParseRpcToken()
		if err != nil {
			return nil, err
		}
		rpcs = append(rpcs, r)
		p.skipWhiteSpaces()
	}
	p.Next()

	return rpcs, nil
}

// rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
func (p *protoParser) ParseRpcToken() (*symbols.Symbol, error) {
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
	if err := p.peekSymbol(';'); err != nil {
		return nil, err
	}
	return s, nil
}

// enum ExampleEnum {
//   ONE = 0;
//   TWO = 1;
//   THREE = 2;
// }
func (p *protoParser) ParseEnumToken() (*symbols.Symbol, error) {
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

	p.skipCurlyBraces()
	return s, nil
}

// message ExampleRPCResponse {}
func (p *protoParser) ParseMessageToken() (*symbols.Symbol, error) {
	s := &symbols.Symbol{}
	s.SetType(token.Message)

	p.skipWhiteSpaces()

	s.SetLine(p.LineNumber())
	s.SetStartChar(p.CharNumber())

	name, err := p.extractName()
	if err != nil {
		return nil, err
	}

	s.SetName(name)
	s.SetEndChar(p.CharNumber())

	p.skipCurlyBraces()
	return s, nil
}
