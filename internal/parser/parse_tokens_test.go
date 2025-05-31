package parser

import (
	"strings"
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func withSpaces(parts ...string) string {
	spaces := "\n \t\n \t"
	return spaces + strings.Join(parts, spaces) + spaces
}

func TestParseTokens_Syntax(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{
			`syntax = "proto3";`,
			`syntax="proto3";`,
			withSpaces("syntax", "=", `"proto3"`, ";"),
		}
		for _, in := range input {
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParseSyntaxToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			// TODO: maybe regexp???
			assert.Equal(t, "syntax", result.Type())
			assert.Equal(t, "proto3", result.Name())
			assert.NoError(t, err)
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
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParseSyntaxToken()

			assert.Nil(t, result)
			assert.Error(t, err)
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
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParsePackageToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "package", result.Type())
			assert.Equal(t, "example", result.Name())
			assert.NoError(t, err)
		}
	})

	t.Run("With errors", func(t *testing.T) {
		input := []string{
			"package example",
			"packageexample",
			"package; example",
		}
		
		for _, in := range input {
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParsePackageToken()

			assert.Nil(t, result)
			assert.Error(t, err)
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
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParseImportToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "import", result.Type())
			assert.Equal(t, "google/protobuf/timestamp.proto", result.Name())
			assert.NoError(t, err)
		}
	})

	t.Run("With Errors", func(t *testing.T) {
		input := []string{
			`import "google/protobuf/timestamp.proto;`,
			`import google/protobuf/timestamp.proto";`,
		}
		
		for _, in := range input {
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParseImportToken()

			assert.Nil(t, result)
			assert.Error(t, err)
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
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParseOptionToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "option", result.Type())
			assert.Equal(t, "go_package", result.Name())
			assert.NoError(t, err)
		}
	})

	t.Run("With Errors", func(t *testing.T) {
		input := []string{
			`option go_package = "gitlab.ozon.ru/example/api/example;example"`,
		}
		
		for _, in := range input {
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParseOptionToken()

			assert.Nil(t, result)
			assert.Error(t, err)
		}
	})
}
