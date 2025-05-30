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
