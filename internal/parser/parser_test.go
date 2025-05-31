package parser

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func TestParser_ExtractKeyword(t *testing.T) {
	input := "hello \t\n\t    world"
	expected := []string{
		"hello",
		"world",
	}
	parser := NewProtoParser(source.NewStringSource(input))

	for _, word := range expected {
		actual, err := parser.extractKeyword()
		assert.Equal(t, word, actual)
		assert.NoError(t, err)
		parser.skipWhiteSpaces()
	}

	actual, err := parser.extractKeyword()
	assert.Equal(t, "", actual)
	assert.Error(t, err)
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

// TODO: test when only one quote, when no quotes, when empty string etc
func TestParser_ExtractQuotedString(t *testing.T) {
	input := "\" hello world \""
	expected := " hello world "
	parser := NewProtoParser(source.NewStringSource(input))

	actual, err := parser.extractQuotedString()
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}
