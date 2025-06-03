package lexer

import (
	"testing"

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

func Test_Lexer_ExtractName(t *testing.T) {
	tests := []testCase{
		{
			"Normal parsing",
			"One_Two_Three",
			"One_Two_Three",
			false,
		}, {
			"Normal parsing",
			"One_Two_Three()",
			"One_Two_Three",
			false,
		},
	}

	runLexerTest(t, (*Lexer).ExtractName, tests)
}

func Test_Lexer_ExtractNameBetweenParens(t *testing.T) {
	tests := []testCase{
		{
			"Normal parsing",
			"(One_Two_Three)",
			"One_Two_Three",
			false,
		}, {
			"No parens",
			"One_Two_Three",
			"",
			true,
		}, {
			"No open paren",
			"One_Two_Three)",
			"",
			true,
		}, {
			"No matching paren",
			"(One_Two_Three",
			"",
			true,
		},
	}

	runLexerTest(t, (*Lexer).ExtractNameBetweenParentheses, tests)
}

func Test_Lexer_SkipCurlyBraces(t *testing.T) {
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
