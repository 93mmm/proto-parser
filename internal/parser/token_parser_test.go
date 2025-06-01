package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func withSpaces(parts ...string) string {
	spaces := "\n \t\n \t"
	return spaces + strings.Join(parts, spaces)
}

func TestParseTokens_Syntax(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			`syntax = "proto3";`,
			`syntax="proto3";`,
			withSpaces("syntax", "=", `"proto3"`, ";"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParseSyntaxToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			// TODO: maybe regexp???
			assert.Equal(t, "syntax", result.Type())
			assert.Equal(t, "proto3", result.Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})

	t.Run("With errors", func(t *testing.T) {
		input := []string{
			`syntax  "proto3";`,
			`syntax = "proto3"`,
			`syntax = "proto3;`,
			`syntax = "proto3` + "\n" + `;`,
		}

		for _, in := range input {
			t.Run(in, func(t *testing.T) {
				parser := newTestTokenParser(in)
				parser.ExtractKeyword()
				result, err := parser.ParseSyntaxToken()

				assert.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})
}

func TestParseTokens_Package(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			"package example;",
			withSpaces("package", "example", ";"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParsePackageToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "package", result.Type())
			assert.Equal(t, "example", result.Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})

	t.Run("With errors", func(t *testing.T) {
		input := []string{
			"package example",
			"packageexample",
			"package; example",
		}

		for _, in := range input {
			t.Run(in, func(t *testing.T) {
				parser := newTestTokenParser(in)
				parser.ExtractKeyword()
				result, err := parser.ParsePackageToken()

				assert.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})
}

func TestParseTokens_Import(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			`import "google/protobuf/timestamp.proto";`,
			withSpaces("import", `"google/protobuf/timestamp.proto"`, ";"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParseImportToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "import", result.Type())
			assert.Equal(t, "google/protobuf/timestamp.proto", result.Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})

	t.Run("With Errors", func(t *testing.T) {
		input := []string{
			`import "google/protobuf/timestamp.proto;`,
			`import google/protobuf/timestamp.proto";`,
		}

		for _, in := range input {
			t.Run(in, func(t *testing.T) {
				parser := newTestTokenParser(in)
				parser.ExtractKeyword()
				result, err := parser.ParseImportToken()

				assert.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})
}

func TestParseTokens_Option(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			`option go_package = "gitlab.ozon.ru/example/api/example;example";`,
			withSpaces("option", "go_package", "=", `"gitlab.ozon.ru/example/api/example;example"`, ";"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParseOptionToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "option", result.Type())
			assert.Equal(t, "go_package", result.Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})

	t.Run("With Errors", func(t *testing.T) {
		input := []string{
			`option go_package = "gitlab.ozon.ru/example/api/example;example"`,
			`option go_package = "gitlab.ozon.ru/example/api/example;example;`,
			`option go_package = gitlab.ozon.ru/example/api/example;example";`,
			`option go_package  "gitlab.ozon.ru/example/api/example;example";`,
			// `optiongo_package = "gitlab.ozon.ru/example/api/example;example";`, // found bug: type=optiongo, name=_package (ignore space absense)
		}

		for _, in := range input {
			t.Run(in, func(t *testing.T) {
				parser := newTestTokenParser(in)
				parser.ExtractKeyword()
				result, err := parser.ParseOptionToken()

				assert.Nil(t, result)
				assert.Error(t, err)
			})
		}
	})
}

func TestParseTokens_Service(t *testing.T) {
	t.Run("Normal, w/o rpcs inside", func(t *testing.T) {
		input := []string{
			`service Example {}`,
			withSpaces("service", "Example", "{", "}"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParseServiceToken()

			assert.Equal(t, 1, len(result))

			assert.Equal(t, "service", result[0].Type())
			assert.Equal(t, "Example", result[0].Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})

	t.Run("Normal, with rpcs inside", func(t *testing.T) {
		input := `service Example {
			rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
			rpc ExampleRPC1(ExampleRPCRequest) returns (ExampleRPCResponse) {};
		}`
		parser := newTestTokenParser(input)
		parser.ExtractKeyword()
		result, err := parser.ParseServiceToken()

		assert.Equal(t, 3, len(result))

		assert.Equal(t, "service", result[0].Type())
		assert.Equal(t, "Example", result[0].Name())

		assert.Equal(t, "rpc", result[1].Type())
		assert.Equal(t, "ExampleRPC", result[1].Name())

		assert.Equal(t, "rpc", result[2].Type())
		assert.Equal(t, "ExampleRPC1", result[2].Name())
		assert.NoError(t, err)
	})
}

func TestParseTokens_Rpc(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			"rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};",
			"rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse);",
			withSpaces("rpc", "ExampleRPC", "(", "ExampleRPCRequest", ")", "returns", "(", "ExampleRPCResponse", ")", "{", "}", ";"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParseRpcToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "rpc", result.Type())
			assert.Equal(t, "ExampleRPC", result.Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})
}

func TestParseTokens_Enum(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			`enum ExampleEnum {
				ONE = 0;
				TWO = 1;
				THREE = 2;
			}`,
			withSpaces("enum", "ExampleEnum", "{", "}"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParseEnumToken()

			assert.Equal(t, "enum", result.Type())
			assert.Equal(t, "ExampleEnum", result.Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})
}

func TestParseTokens_Message(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			`message ExampleRPCResponse {}`,
			`message ExampleRPCResponse {
				message Emb { string field11 = 1; }
				ExampleEnum field1 = 1;
				Emb filed2 = 2;
				google.protobuf.Timestamp filed3 = 3;
			}`,
			withSpaces("message", "ExampleRPCResponse", "{", "}"),
		}
		for _, in := range input {
			parser := newTestTokenParser(in)
			parser.ExtractKeyword()
			result, err := parser.ParseMessageToken()

			assert.Equal(t, "message", result.Type())
			assert.Equal(t, "ExampleRPCResponse", result.Name())
			assert.NoError(t, err)
			assert.True(t, parser.EOF())
		}
	})
}
