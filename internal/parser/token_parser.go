package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/errors"
	"github.com/93mmm/proto-parser/internal/lexer"
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

type tokenParser struct {
	*lexer.Lexer
}

func NewTokenParser(l *lexer.Lexer) *tokenParser {
	return &tokenParser{
		Lexer: l,
	}
}

func newTestTokenParser(in string) *tokenParser {
	src := source.NewStringSource(in)
	bp := base.NewBaseParser(src)
	l := lexer.NewLexer(bp)
	return NewTokenParser(l)
}

// syntax = "proto3";
func (p *tokenParser) ParseSyntaxToken() (*symbols.Symbol, error) {
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
func (p *tokenParser) ParsePackageToken() (*symbols.Symbol, error) {
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
func (p *tokenParser) ParseImportToken() (*symbols.Symbol, error) {
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
func (p *tokenParser) ParseOptionToken() (*symbols.Symbol, error) {
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
func (p *tokenParser) ParseServiceToken() ([]*symbols.Symbol, error) {
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
			return nil, errors.NewLexerError(p.LineNumber(), p.CharNumber(), "Unexpected keyword %s found inside %s", keyword, s)
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
func (p *tokenParser) ParseRpcToken() (*symbols.Symbol, error) {
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
func (p *tokenParser) ParseEnumToken() (*symbols.Symbol, error) {
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
func (p *tokenParser) ParseMessageToken() (*symbols.Symbol, error) {
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
