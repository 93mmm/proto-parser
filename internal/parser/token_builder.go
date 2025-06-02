package parser

import (
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
)

var tokenBuilders = map[string]TokenBuilder{
	token.Syntax:  SyntaxToken{},
	token.Package: PackageToken{},
	token.Import:  ImportToken{},
	token.Option:  OptionToken{},
	token.Service: ServiceToken{},
	token.Enum:    EnumToken{},
	token.Message: MessageToken{},
}

func getBuilder(token string) (TokenBuilder, bool) {
	b, ok := tokenBuilders[token]
	return b, ok
}

type TokenBuilder interface {
	Parse(*tokenParser, symbols.Collector) error
}

type SyntaxToken struct{}

// func (_ SyntaxToken) Parse(parser *tokenParser, collector symbols.Collector) error {
// 	return parser.ParseSyntaxToken(collector)
// }

// TODO: i want to rewrite this
func (_ SyntaxToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	s, err := parser.ParseSyntaxToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type PackageToken struct{}

func (_ PackageToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	s, err := parser.ParsePackageToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type ImportToken struct{}

func (_ ImportToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	s, err := parser.ParseImportToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type OptionToken struct{}

func (_ OptionToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	s, err := parser.ParseOptionToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type ServiceToken struct{}

func (_ ServiceToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	s, err := parser.ParseServiceToken()
	if err != nil {
		return err
	}
	collector.Add(s...)
	return nil
}

type EnumToken struct{}

func (_ EnumToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	s, err := parser.ParseEnumToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type MessageToken struct{}

func (_ MessageToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	s, err := parser.ParseMessageToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}
