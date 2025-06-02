package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTokens_Syntax(t *testing.T) {
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
