package builder

import (
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/parser/constants"
)

var tokenBuilders = map[string]TokenBuilder{
	constants.Syntax:  SyntaxToken{},
	constants.Package: PackageToken{},
	constants.Import:  ImportToken{},
	constants.Option:  OptionToken{},
	constants.Service: ServiceToken{},
	constants.Enum:    EnumToken{},
	constants.Message: MessageToken{},
}

func GetBuilder(token string) (TokenBuilder, bool) {
	b, ok := tokenBuilders[token]
	return b, ok
}

type TokenBuilder interface {
	Parse(*TokenParser, symbols.Collector) error
}

type SyntaxToken struct{}

func (_ SyntaxToken) Parse(parser *TokenParser, collector symbols.Collector) error {
	s, err := parser.ParseSyntaxToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type PackageToken struct{}

func (_ PackageToken) Parse(parser *TokenParser, collector symbols.Collector) error {
	s, err := parser.ParsePackageToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type ImportToken struct{}

func (_ ImportToken) Parse(parser *TokenParser, collector symbols.Collector) error {
	s, err := parser.ParseImportToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type OptionToken struct{}

func (_ OptionToken) Parse(parser *TokenParser, collector symbols.Collector) error {
	s, err := parser.ParseOptionToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type ServiceToken struct{}

func (_ ServiceToken) Parse(parser *TokenParser, collector symbols.Collector) error {
	s, err := parser.ParseServiceToken()
	if err != nil {
		return err
	}
	collector.Add(s...)
	return nil
}

type EnumToken struct{}

func (_ EnumToken) Parse(parser *TokenParser, collector symbols.Collector) error {
	s, err := parser.ParseEnumToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}

type MessageToken struct{}

func (_ MessageToken) Parse(parser *TokenParser, collector symbols.Collector) error {
	s, err := parser.ParseMessageToken()
	if err != nil {
		return err
	}
	collector.Add(s)
	return nil
}
