package tests

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/parser"
	"github.com/93mmm/proto-parser/internal/parser/base"
	"github.com/93mmm/proto-parser/internal/parser/builder"
	"github.com/93mmm/proto-parser/internal/parser/lexer"
	"github.com/93mmm/proto-parser/internal/parser/source"
	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/parser/constants"
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
		{Name: "proto3", Type: constants.Syntax, Line: 1, Start: 10, End: 18},
		{Name: "example", Type: constants.Package, Line: 3, Start: 9, End: 16},
		{Name: "google/protobuf/timestamp.proto", Type: constants.Import, Line: 5, Start: 8, End: 41},
		{Name: "go_package", Type: constants.Option, Line: 7, Start: 8, End: 18},
		{Name: "Example", Type: constants.Service, Line: 9, Start: 9, End: 16},
		{Name: "ExampleRPC", Type: constants.Rpc, Line: 10, Start: 7, End: 17},
		{Name: "ExampleEnum", Type: constants.Enum, Line: 13, Start: 6, End: 17},
		{Name: "ExampleRPCRequest", Type: constants.Message, Line: 19, Start: 9, End: 26},
		{Name: "ExampleRPCResponse", Type: constants.Message, Line: 26, Start: 9, End: 27},
	})
}

func Test_Integration_Syntax(t *testing.T) {
	runTestFromFile(t, "testdata/syntax.proto", []*symbols.Symbol{
		{Name: "proto3", Type: constants.Syntax, Line: 1, Start: 10, End: 18},
	})
}

func Test_Integration_Package(t *testing.T) {
	runTestFromFile(t, "testdata/package.proto", []*symbols.Symbol{
		{Name: "example", Type: constants.Package, Line: 1, Start: 9, End: 16},
	})
}

func Test_Integration_Import(t *testing.T) {
	runTestFromFile(t, "testdata/import.proto", []*symbols.Symbol{
		{Name: "google/protobuf/timestamp.proto", Type: constants.Import, Line: 1, Start: 8, End: 41},
	})
}

func Test_Integration_Option(t *testing.T) {
	runTestFromFile(t, "testdata/option.proto", []*symbols.Symbol{
		{Name: "go_package", Type: constants.Option, Line: 1, Start: 8, End: 18},
	})
}

func Test_Integration_Service(t *testing.T) {
	runTestFromFile(t, "testdata/service.proto", []*symbols.Symbol{
		{Name: "Example", Type: constants.Service, Line: 1, Start: 9, End: 16},
		{Name: "Example", Type: constants.Service, Line: 3, Start: 9, End: 16},
		{Name: "ExampleRPC", Type: constants.Rpc, Line: 4, Start: 7, End: 17},
		{Name: "ExampleRPC1", Type: constants.Rpc, Line: 5, Start: 7, End: 18},
		{Name: "ExampleRPC2", Type: constants.Rpc, Line: 6, Start: 7, End: 18},
		{Name: "ExampleRPC3", Type: constants.Rpc, Line: 7, Start: 7, End: 18},
		{Name: "ExampleRPC4", Type: constants.Rpc, Line: 8, Start: 7, End: 18},
	})
}

func Test_Integration_Enum(t *testing.T) {
	runTestFromFile(t, "testdata/enum.proto", []*symbols.Symbol{
		{Name: "ExampleEnum", Type: constants.Enum, Line: 1, Start: 6, End: 17},
		{Name: "ExampleEnum", Type: constants.Enum, Line: 2, Start: 6, End: 17},
	})
}

func Test_Integration_Message(t *testing.T) {
	runTestFromFile(t, "testdata/message.proto", []*symbols.Symbol{
		{Name: "ExampleRPCRequest", Type: constants.Message, Line: 1, Start: 9, End: 26},
		{Name: "ExampleRPCRequest", Type: constants.Message, Line: 2, Start: 9, End: 26},
		{Name: "ExampleRPCRequest", Type: constants.Message, Line: 8, Start: 9, End: 26},
	})
}
