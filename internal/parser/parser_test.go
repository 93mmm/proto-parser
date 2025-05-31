package parser

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func TestParser_ExtractKeyword(t *testing.T) {
	t.Run("Regular string", func(t *testing.T) {
		input := "one two three\n\t\n four five!"
		parser := NewProtoParser(source.NewStringSource(input))
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
		parser := NewProtoParser(source.NewStringSource(input))

		actual, err := parser.extractKeyword()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Not keyword", func(t *testing.T) {
		input := ",hello"
		expected := ""
		parser := NewProtoParser(source.NewStringSource(input))

		actual, err := parser.extractKeyword()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})
}

// TODO: test when only one quote, when no quotes, when empty string etc
func TestParser_ExtractQuotedString(t *testing.T) {
	t.Run("Normal quotes", func(t *testing.T) {
		input := "\" one \""
		expected := " one "
		parser := NewProtoParser(source.NewStringSource(input))

		actual, err := parser.extractQuotedString()
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Quote w/o pair", func(t *testing.T) {
		input := "\" one "
		expected := ""
		parser := NewProtoParser(source.NewStringSource(input))

		actual, err := parser.extractQuotedString()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})

	t.Run("Quote with pair on next string", func(t *testing.T) {
		input := "\" one \n\""
		expected := ""
		parser := NewProtoParser(source.NewStringSource(input))

		actual, err := parser.extractQuotedString()
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
	})
}

func TestParser_SkipUntilNextLine(t *testing.T) {
	input := "hello world this is\n fine!"
	parser := NewProtoParser(source.NewStringSource(input))

	parser.extractKeyword()
	parser.skipUntilNextLine()
	parser.skipWhiteSpaces()

	actual, err := parser.extractKeyword()
	assert.Equal(t, "fine", actual)
	assert.NoError(t, err)
}
