package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/errors"
	"github.com/93mmm/proto-parser/internal/lexer"
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
)

type tokenParser struct {
	factory symbols.SymbolFactory
	*lexer.Lexer
}

func NewTokenParser(l *lexer.Lexer) *tokenParser {
	return &tokenParser{
		Lexer: l,
		factory: symbols.NewSymbolFactory(),
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
	if err := p.PeekSymbol('='); err != nil {
		return nil, err
	}

	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()
	name, err := p.ExtractQuotedString()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()

	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}

	s := p.factory.NewSyntaxSymbol(
		name, line, start, end,
	)
	return s, nil
}

// package example;
func (p *tokenParser) ParsePackageToken() (*symbols.Symbol, error) {
	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()
	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()
	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}
	s := p.factory.NewPackageSymbol(
		name, line, start, end,
	)
	return s, nil
}

// import "google/protobuf/timestamp.proto";
func (p *tokenParser) ParseImportToken() (*symbols.Symbol, error) {
	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()
	name, err := p.ExtractQuotedString()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()
	if err := p.PeekSymbol(';'); err != nil {
		return nil, err
	}
	s := p.factory.NewImportSymbol(
		name, line, start, end,
	)
	return s, nil
}

// option go_package = "gitlab.ozon.ru/example/api/example;example";
func (p *tokenParser) ParseOptionToken() (*symbols.Symbol, error) {
	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()
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
	s := p.factory.NewOptionSymbol(
		name, line, start, end,
	)
	return s, nil
}

//	service Example {
//	  rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
//	}
func (p *tokenParser) ParseServiceToken() ([]*symbols.Symbol, error) {
	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()
	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()
	p.SkipWhiteSpaces()

	if err := p.PeekSymbol('{'); err != nil {
		return nil, err
	}

	p.SkipWhiteSpaces()

	s := p.factory.NewServiceSymbol(
		name, line, start, end,
	)
	result := make([]*symbols.Symbol, 0, 10)
	result = append(result, s)
	for !p.Test('}') && !p.EOF() {
		keyword, _ := p.ExtractKeyword()
		if keyword != "rpc" {
			return result, errors.NewError(p.LineNumber(), p.CharNumber(), "Unexpected keyword %s found inside %s", keyword, s)
		}
		s, err := p.ParseRpcToken()
		if err != nil {
			return result, err
		}
		result = append(result, s)
		p.SkipWhiteSpaces()
	}
	p.Next()

	return result, nil
}

// rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
func (p *tokenParser) ParseRpcToken() (*symbols.Symbol, error) {
	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()
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
	s := p.factory.NewRpcSymbol(
		name, line, start, end,
	)
	return s, nil
}

//	enum ExampleEnum {
//	  ONE = 0;
//	  TWO = 1;
//	  THREE = 2;
//	}
func (p *tokenParser) ParseEnumToken() (*symbols.Symbol, error) {
	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()
	p.SkipCurlyBraces()

	s := p.factory.NewEnumSymbol(
		name, line, start, end,
	)
	return s, nil
}

// message ExampleRPCResponse {}
func (p *tokenParser) ParseMessageToken() (*symbols.Symbol, error) {
	p.SkipWhiteSpaces()

	line, start := p.LineNumber(), p.CharNumber()

	name, err := p.ExtractName()
	if err != nil {
		return nil, err
	}

	end := p.CharNumber()
	p.SkipCurlyBraces()

	s := p.factory.NewMessageSymbol(
		name, line, start, end,
	)
	return s, nil
}
