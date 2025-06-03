package symbols

import "github.com/93mmm/proto-parser/internal/parser/constants"

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

type defaultSymbolFactory struct{}

func NewSymbolFactory() *defaultSymbolFactory {
	return &defaultSymbolFactory{}
}

func (_ defaultSymbolFactory) NewSyntaxSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Syntax, line, start, end,
	)
}

func (_ defaultSymbolFactory) NewPackageSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Package, line, start, end,
	)
}

func (_ defaultSymbolFactory) NewImportSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Import, line, start, end,
	)
}

func (_ defaultSymbolFactory) NewOptionSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Option, line, start, end,
	)
}

func (_ defaultSymbolFactory) NewServiceSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Service, line, start, end,
	)
}

func (_ defaultSymbolFactory) NewRpcSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Rpc, line, start, end,
	)
}

func (_ defaultSymbolFactory) NewEnumSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Enum, line, start, end,
	)
}

func (_ defaultSymbolFactory) NewMessageSymbol(name string, line, start, end int) *Symbol {
	return NewSymbol(
		name, constants.Message, line, start, end,
	)
}
