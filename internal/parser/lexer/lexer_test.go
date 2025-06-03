package lexer

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/parser/constants"
	"github.com/stretchr/testify/assert"
)

func TestParser_ExtractKeyword(t *testing.T) {
	t.Run("Regular string", func(t *testing.T) {
		input := "one two three\n\t\n four five!"
		parser := newTestLexer(input)
		expected := []string{
			"one", "two", "three", "four", "five",
		}

		for _, e := range expected {
			actual, err := parser.ExtractKeyword()
			assert.Equal(t, e, actual)
			assert.NoError(t, err)
			parser.SkipWhiteSpaces()
		}
		actual, err := parser.ExtractKeyword()
		assert.Equal(t, "", actual)
		assert.Error(t, err)
	})

	t.Run("Empty string", func(t *testing.T) {
		input := ""
		expected := ""
		parser := newTestLexer(input)

		actual, err := parser.ExtractKeyword()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Not keyword", func(t *testing.T) {
		input := ",hello"
		expected := ""
		parser := newTestLexer(input)

		actual, err := parser.ExtractKeyword()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Not keyword", func(t *testing.T) { // TODO: edit
		input := "hello_world"
		expected := "hello"
		parser := newTestLexer(input)

		actual, err := parser.ExtractKeyword()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}

func TestParser_ExtractQuotedString(t *testing.T) {
	t.Run("Normal quotes", func(t *testing.T) {
		input := "\" one \""
		expected := " one "
		parser := newTestLexer(input)

		actual, err := parser.ExtractQuotedString()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Quote w/o pair", func(t *testing.T) {
		input := "\" one "
		expected := ""
		parser := newTestLexer(input)

		actual, err := parser.ExtractQuotedString()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Quote with pair on next string", func(t *testing.T) {
		input := "\" one \n\""
		expected := ""
		parser := newTestLexer(input)

		actual, err := parser.ExtractQuotedString()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})
}

func TestParser_ExtractName(t *testing.T) {
	t.Run("Normal names", func(t *testing.T) {
		input := "One_Two_Three"
		expected := "One_Two_Three"
		parser := newTestLexer(input)

		actual, err := parser.ExtractName()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}

func TestParser_ExtractNameBetweenParantheses(t *testing.T) {
	t.Run("Normal names", func(t *testing.T) {
		input := "(One_Two_Three)"
		expected := "One_Two_Three"
		parser := newTestLexer(input)

		actual, err := parser.ExtractNameBetweenParentheses()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)

		actual, err = parser.ExtractKeyword()
		assert.Equal(t, "", actual)
		assert.Error(t, err)
	})
}

func TestParser_SkipUntilNextLine(t *testing.T) {
	input := "one two : three"
	parser := newTestLexer(input)

	parser.ExtractKeyword()
	parser.SkipUntilMatch(constants.Colon)

	assert.True(t, parser.Test(constants.Colon))

	parser.Next()

	actual, err := parser.ExtractKeyword()
	assert.Equal(t, "three", actual)
	assert.NoError(t, err)
}

func TestParser_SkipCurlyBraces(t *testing.T) {
	input := "{{{{{{{{{{{{{{}}}}}}}}}}}}}}"
	parser := newTestLexer(input)
	parser.SkipCurlyBraces()

	t.Log(parser.CharNumber())
	assert.True(t, parser.EOF())
}

func Test_Parser_SkipCurlyBraces(t *testing.T) {
	type result struct {
		ret bool
		eof bool
	}
	tests := []struct {
		name  string
		input string
		want  result
	}{
		{
			"Ok",
			"{}",
			result{true, true},
		}, {
			"Ok",
			"{{{{{}}}}}",
			result{true, true},
		}, {
			"Ok",
			"{{{{{}}}}}         ",
			result{true, false},
		}, {
			"Not ok",
			"{",
			result{false, true},
		}, {
			"Not ok",
			"{{{}}",
			result{false, true},
		}, {
			"Not ok",
			"{{{}}           ",
			result{false, true},
		}, {
			"Not ok",
			"{{{}}           ",
			result{false, true},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lexer := newTestLexer(test.input)
			t.Log(lexer.CharNumber())
			
			assert.Equal(t, test.want.ret, lexer.SkipCurlyBraces())
			assert.Equal(t, test.want.eof, lexer.EOF())
		})
	}
}
