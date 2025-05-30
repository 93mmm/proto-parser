package baseparser

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func TestBaseParser(t *testing.T) {
	input := "hello world"

	t.Run("Next", func(t *testing.T) {
		parser := NewBaseParser(
			source.NewStringSource(input),
		)
		for _, c := range input {
			assert.Equal(t, c, parser.Next())
		}
	})

	t.Run("Peek", func(t *testing.T) {
		parser := NewBaseParser(
			source.NewStringSource(input),
		)
		for _, c := range input {
			negative := c + 1
			assert.False(t, parser.Peek(negative))
			assert.True(t, parser.Peek(c))
		}
	})

	t.Run("EOF", func(t *testing.T) {
		parser := NewBaseParser(
			source.NewStringSource(input),
		)
		checkLength := len([]rune(input))
		ptr := 0

		for !parser.EOF() {
			ptr++
			parser.Next()
			if !assert.GreaterOrEqual(t, checkLength, ptr) {
				t.Error("Wrong EOF method!")
				break
			}
		}
		assert.True(t, parser.EOF())
	})
}

func TestBaseParser_CurrentLineAndChar(t *testing.T) {
	input := "ttt\ni\nl"
	type pos struct {
		char int
		line int
	}
	positions := []pos{
		{1, 1},
		{2, 1},
		{3, 1},
		{0, 2},
		{1, 2},
		{0, 3},
		{1, 3},
		// end of file reached
		{1, 3},
		{1, 3},
		{1, 3},
	}
	parser := NewBaseParser(
		source.NewStringSource(input),
	)

	for _, p := range positions {
		assert.Equal(t, p.char, parser.CharNumber())
		assert.Equal(t, p.line, parser.LineNumber())
		parser.Next()
	}
}
