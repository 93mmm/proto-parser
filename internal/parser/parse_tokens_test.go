package parser

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func TestParseTokens_Syntax(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		spaces := "\n\t\n\t"
		input := []string{
			`syntax = "proto3";`,
			`syntax="proto3";`,
			spaces + `syntax="proto3";` + spaces,
			spaces + `syntax` + spaces + `=` + spaces + `"proto3"` + spaces + `;`,
		}
		for _, in := range input {
			parser := NewProtoParser(source.NewStringSource(in))
			parser.extractKeyword()
			result, err := parser.ParseSyntaxToken()

			if result == nil {
				t.Error("result is nil", err)
				continue
			}
			assert.Equal(t, "syntax", result.Type())
			assert.Equal(t, "proto3", result.Name())
			assert.NoError(t, err)
		}
	})

	t.Run("With errors", func(t *testing.T) {
		// input := `syntax = "proto3";`
		// result, err := NewProtoParser(source.NewStringSource(input)).ParseSyntaxToken()
		//
		// assert.Equal(t, "syntax", result.Type())
		// assert.Equal(t, "proto3", result.Name())
		// assert.NoError(t, err)
	})
}

func TestParseTokens_Package(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		spaces := "\n\t\n\t"
		input := []string{
			"package example;",
			spaces + "package" + spaces + "example" + spaces + ";" + spaces,
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
}
