package tests

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/parser"
	"github.com/93mmm/proto-parser/internal/parser/base"
	"github.com/93mmm/proto-parser/internal/parser/builder"
	"github.com/93mmm/proto-parser/internal/parser/lexer"
	"github.com/93mmm/proto-parser/internal/parser/source"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
	"github.com/stretchr/testify/assert"
)

func runTestFromFile(t *testing.T, file string, expected []*symbols.Symbol) {
	src, err := source.NewFileSource(file)
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()
	bp := base.NewBaseParser(src)
	l := lexer.NewLexer(bp)
	pp := builder.NewTokenParser(l)
	parsed, err := parser.NewParser(pp).ParseDocument()

	assert.NoError(t, err)
	if !assert.Equal(t, len(expected), len(parsed)) {
		t.Fatal("Parsed length is diffrent from expected")
	}

	for i, s := range expected {
		assert.Equal(t, *s, *parsed[i])
	}
}

func Test_Integration_Full(t *testing.T) {
	runTestFromFile(t, "testdata/full.proto", []*symbols.Symbol{
		{Name: "proto3", Type: token.Syntax, Line: 1, Start: 10, End: 18},
		{Name: "example", Type: token.Package, Line: 3, Start: 9, End: 16},
		{Name: "google/protobuf/timestamp.proto", Type: token.Import, Line: 5, Start: 8, End: 41},
		{Name: "go_package", Type: token.Option, Line: 7, Start: 8, End: 18},
		{Name: "Example", Type: token.Service, Line: 9, Start: 9, End: 16},
		{Name: "ExampleRPC", Type: token.Rpc, Line: 10, Start: 7, End: 17},
		{Name: "ExampleEnum", Type: token.Enum, Line: 13, Start: 6, End: 17},
		{Name: "ExampleRPCRequest", Type: token.Message, Line: 19, Start: 9, End: 26},
		{Name: "ExampleRPCResponse", Type: token.Message, Line: 26, Start: 9, End: 27},
	})
}

func Test_Integration_Syntax(t *testing.T) {
	runTestFromFile(t, "testdata/syntax.proto", []*symbols.Symbol{
		{Name: "proto3", Type: token.Syntax, Line: 1, Start: 10, End: 18},
	})
}

func Test_Integration_Package(t *testing.T) {
	runTestFromFile(t, "testdata/package.proto", []*symbols.Symbol{
		{Name: "example", Type: token.Package, Line: 1, Start: 9, End: 16},
	})
}
