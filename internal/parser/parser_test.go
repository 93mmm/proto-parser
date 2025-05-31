package parser

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func TestParser_ExtractWord(t *testing.T) {
	input := "hello \t\n\t    world"
	expected := []string{
		"hello",
		"world",
	}
	parser := NewProtoParser(source.NewStringSource(input))

	for _, word := range expected {
		assert.Equal(t, word, parser.extractKeyword())
		parser.skipWhiteSpaces()
	}
	assert.Equal(t, "", parser.extractKeyword())
}

func TestParser_SkipUntilNextLine(t *testing.T) {
	input := "hello world this is\n fine!"
	parser := NewProtoParser(source.NewStringSource(input))

	parser.extractKeyword()
	parser.skipUntilNextLine()
	parser.skipWhiteSpaces()

	assert.Equal(t, "fine", parser.extractKeyword())
}

