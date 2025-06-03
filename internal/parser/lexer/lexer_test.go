package lexer

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/parser/constants"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name  string
	input string
	want  string
	err   bool
}

type extractX func(*Lexer) (string, error)

func runLexerTest(t *testing.T, extract extractX, tests []testCase) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lexer := newTestLexer(test.input)
			actual, err := extract(lexer)

			assert.Equal(t, test.want, actual)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_Lexer_ExtractKeyword(t *testing.T) {
	tests := []testCase{
		{
			"Normal parsing",
			"input",
			"input",
			false,
		}, {
			"Empty string",
			"",
			"",
			true,
		}, {
			"Invalid start",
			",input",
			"",
			true,
		}, {
			"Invalid keyword camelCase",
			"helloEveryone",
			"",
			true,
		}, {
			"Invalid keyword PascalCase",
			"HelloEveryone",
			"",
			true,
		}, {
			"Invalid keyword snake_case",
			"hello_everyone",
			"",
			true,
		}, {
			"Invalid keyword number in between",
			"hello4everyone",
			"",
			true,
		},
	}
	runLexerTest(t, (*Lexer).ExtractKeyword, tests)
}

func Test_Lexer_ExtractQuotedString(t *testing.T) {
	tests := []testCase{
		{
			"Normal parsing",
			"\"input\"",
			"input",
			false,
		}, {
			"Quote w/o pair",
			"\"input",
			"",
			true,
		}, {
			"Without quote in begin",
			"input\n\"",
			"",
			true,
		}, {
			"\\n between quotes",
			"\"input\n\"",
			"",
			true,
		},
	}

	runLexerTest(t, (*Lexer).ExtractQuotedString, tests)
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
