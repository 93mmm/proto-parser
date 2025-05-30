package parser

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func TestBaseParser(t *testing.T) {
	input := "hello \t\n\t    world"
	expected := []string{
		"hello",
		"world",
	}
	parser := NewProtoParser(source.NewStringSource(input))

	for _, word := range expected {
		assert.Equal(t, word, parser.extractWord())
		parser.skipWhiteSpaces()
	}
	assert.Equal(t, "", parser.extractWord())
}
