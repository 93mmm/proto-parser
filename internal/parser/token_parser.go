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

type SyntaxToken struct{}

func (_ SyntaxToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseSyntaxToken(collector)
}

// syntax = "proto3";
func (p *tokenParser) ParseSyntaxToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Syntax)

	if err := p.PeekSymbol('='); err != nil {
		return err
	}

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractQuotedString()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol(';'); err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type PackageToken struct{}

func (_ PackageToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParsePackageToken(collector)
}

// package example;
func (p *tokenParser) ParsePackageToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Package)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol(';'); err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type ImportToken struct{}

func (_ ImportToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseImportToken(collector)
}

// import "google/protobuf/timestamp.proto";
func (p *tokenParser) ParseImportToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Import)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractQuotedString()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol(';'); err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type OptionToken struct{}

func (_ OptionToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseOptionToken(collector)
}

// option go_package = "gitlab.ozon.ru/example/api/example;example";
func (p *tokenParser) ParseOptionToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Option)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if err := p.PeekSymbol('='); err != nil {
		return err
	}

	p.SkipWhiteSpaces()
	if _, err := p.ExtractQuotedString(); err != nil {
		return err
	}

	if err := p.PeekSymbol(';'); err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type ServiceToken struct{}

func (_ ServiceToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseServiceToken(collector)
}

//	service Example {
//	  rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
//	}
func (p *tokenParser) ParseServiceToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Service)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	p.SkipWhiteSpaces()

	if err := p.PeekSymbol('{'); err != nil {
		return err
	}

	p.SkipWhiteSpaces()
	collector.Add(s)
	for !p.Test('}') && !p.EOF() {
		keyword, _ := p.ExtractKeyword()
		if keyword != "rpc" {
			return errors.NewError(p.LineNumber(), p.CharNumber(), "Unexpected keyword %s found inside %s", keyword, s)
		}
		err := p.ParseRpcToken(collector)
		if err != nil {
			return err
		}
		p.SkipWhiteSpaces()
	}
	p.Next()

	return nil
}

// rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
func (p *tokenParser) ParseRpcToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Rpc)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	if _, err := p.ExtractNameBetweenParentheses(); err != nil {
		return err
	}
	p.ExtractKeyword()
	if _, err := p.ExtractNameBetweenParentheses(); err != nil {
		return err
	}

	p.SkipUntilMatch(';')
	if err := p.PeekSymbol(';'); err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type EnumToken struct{}

func (_ EnumToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseEnumToken(collector)
}

//	enum ExampleEnum {
//	  ONE = 0;
//	  TWO = 1;
//	  THREE = 2;
//	}
func (p *tokenParser) ParseEnumToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Enum)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	p.SkipCurlyBraces()
	collector.Add(s)
	return nil
}

type MessageToken struct{}

func (_ MessageToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseMessageToken(collector)
}

// message ExampleRPCResponse {}
func (p *tokenParser) ParseMessageToken(collector symbols.Collector) error {
	s := &symbols.Symbol{}
	s.SetType(token.Message)

	p.SkipWhiteSpaces()

	s.SetLine(p.LineNumber()).
		SetStartChar(p.CharNumber())

	name, err := p.ExtractName()
	if err != nil {
		return err
	}

	s.SetName(name).
		SetEndChar(p.CharNumber())

	p.SkipCurlyBraces()
	collector.Add(s)
	return nil
}
