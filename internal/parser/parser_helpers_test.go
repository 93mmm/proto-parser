package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ExtractKeyword(t *testing.T) {
	t.Run("Regular string", func(t *testing.T) {
		input := "one two three\n\t\n four five!"
		parser := newTestProtoParser(input)
		expected := []string{
			"one", "two", "three", "four", "five",
		}

		for _, e := range expected {
			actual, err := parser.extractKeyword()
			assert.Equal(t, e, actual)
			assert.NoError(t, err)
			parser.skipWhiteSpaces()
		}
		actual, err := parser.extractKeyword()
		assert.Equal(t, "", actual)
		assert.Error(t, err)
	})

	t.Run("Empty string", func(t *testing.T) {
		input := ""
		expected := ""
		parser := newTestProtoParser(input)

		actual, err := parser.extractKeyword()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Not keyword", func(t *testing.T) {
		input := ",hello"
		expected := ""
		parser := newTestProtoParser(input)

		actual, err := parser.extractKeyword()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Not keyword", func(t *testing.T) { // TODO: edit
		input := "hello_world"
		expected := "hello"
		parser := newTestProtoParser(input)

		actual, err := parser.extractKeyword()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}

func TestParser_ExtractQuotedString(t *testing.T) {
	t.Run("Normal quotes", func(t *testing.T) {
		input := "\" one \""
		expected := " one "
		parser := newTestProtoParser(input)

		actual, err := parser.extractQuotedString()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Quote w/o pair", func(t *testing.T) {
		input := "\" one "
		expected := ""
		parser := newTestProtoParser(input)

		actual, err := parser.extractQuotedString()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Quote with pair on next string", func(t *testing.T) {
		input := "\" one \n\""
		expected := ""
		parser := newTestProtoParser(input)

		actual, err := parser.extractQuotedString()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})
}

func TestParser_ExtractName(t *testing.T) {
	t.Run("Normal names", func(t *testing.T) {
		input := "One_Two_Three"
		expected := "One_Two_Three"
		parser := newTestProtoParser(input)

		actual, err := parser.extractName()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}

func TestParser_ExtractNameBetweenParantheses(t *testing.T) {
	t.Run("Normal names", func(t *testing.T) {
		input := "(One_Two_Three)"
		expected := "One_Two_Three"
		parser := newTestProtoParser(input)

		actual, err := parser.extractNameBetweenParentheses()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)

		actual, err = parser.extractKeyword()
		assert.Equal(t, "", actual)
		assert.Error(t, err)
	})
}

func TestParser_SkipUntilNextLine(t *testing.T) {
	input := "one two : three"
	parser := newTestProtoParser(input)

	parser.extractKeyword()
	parser.skipUntilMatch(':')

	assert.True(t, parser.Test(':'))

	parser.Next()

	actual, err := parser.extractKeyword()
	assert.Equal(t, "three", actual)
	assert.NoError(t, err)
}

func TestParser_SkipCurlyBraces(t *testing.T) {
	input := "{{{{{{{{{{{{{{}}}}}}}}}}}}}}"
	parser := newTestProtoParser(input)
	parser.skipCurlyBraces()

	t.Log(parser.CharNumber())
	assert.True(t, parser.EOF())
}
