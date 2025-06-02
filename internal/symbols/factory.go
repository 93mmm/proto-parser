package symbols

import "github.com/93mmm/proto-parser/internal/token"

type SymbolFactory interface {
	NewSyntaxSymbol(name string, line, start, end int) *Symbol
	NewPackageSymbol(name string, line, start, end int) *Symbol
	NewImportSymbol(name string, line, start, end int) *Symbol
	NewOptionSymbol(name string, line, start, end int) *Symbol
	NewServiceSymbol(name string, line, start, end int) *Symbol
	NewRpcSymbol(name string, line, start, end int) *Symbol
	NewEnumSymbol(name string, line, start, end int) *Symbol
	NewMessageSymbol(name string, line, start, end int) *Symbol
}

type DefaultSymbolFactory struct{}

func (_ DefaultSymbolFactory) NewSyntaxSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Syntax, line, start, end,
	)
}

func (_ DefaultSymbolFactory) NewPackageSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Package, line, start, end,
	)
}

func (_ DefaultSymbolFactory) NewImportSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Import, line, start, end,
	)
}

func (_ DefaultSymbolFactory) NewOptionSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Option, line, start, end,
	)
}

func (_ DefaultSymbolFactory) NewServiceSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Service, line, start, end,
	)
}

func (_ DefaultSymbolFactory) NewRpcSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Rpc, line, start, end,
	)
}

func (_ DefaultSymbolFactory) NewEnumSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Enum, line, start, end,
	)
}

func (_ DefaultSymbolFactory) NewMessageSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, token.Message, line, start, end,
	)
}
