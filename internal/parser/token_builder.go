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

func (_ SyntaxToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseSyntaxToken(collector)
}

type PackageToken struct{}

func (_ PackageToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParsePackageToken(collector)
}

type ImportToken struct{}

func (_ ImportToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseImportToken(collector)
}

type OptionToken struct{}

func (_ OptionToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseOptionToken(collector)
}

type ServiceToken struct{}

func (_ ServiceToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseServiceToken(collector)
}

type EnumToken struct{}

func (_ EnumToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseEnumToken(collector)
}

type MessageToken struct{}

func (_ MessageToken) Parse(parser *tokenParser, collector symbols.Collector) error {
	return parser.ParseMessageToken(collector)
}
