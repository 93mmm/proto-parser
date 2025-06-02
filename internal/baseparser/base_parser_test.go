package baseparser

import (
	"testing"

	"github.com/93mmm/proto-parser/internal/source"
	"github.com/stretchr/testify/assert"
)

func Test_Next(t *testing.T) {
	input := "test string"
	parser := NewBaseParser(
		source.NewStringSource(input),
	)
	for _, c := range input {
		assert.Equal(t, c, parser.Next())
	}
}

func Test_Peek(t *testing.T) {
	input := "test string"

	parser := NewBaseParser(
		source.NewStringSource(input),
	)
	for _, c := range input {
		wrong := c + 1
		assert.False(t, parser.Peek(wrong))
		assert.True(t, parser.Peek(c))
	}
}

func Test_EOF(t *testing.T) {
	input := "test string"

	parser := NewBaseParser(
		source.NewStringSource(input),
	)
	checkLength := len([]rune(input))
	ptr := 0

	for !parser.EOF() {
		ptr++
		parser.Next()
		if !assert.GreaterOrEqual(t, checkLength, ptr) {
			t.Error("Wrong EOF method")
			break
		}
	}
	assert.True(t, parser.EOF())
}

func Test_CurrentLineAndChar(t *testing.T) {
	type pos struct {
		char int
		line int
	}
	tests := []struct {
		name  string
		input string
		want  []pos
	}{
		{
			name:  "Non-empty",
			input: "ttt\ni\nl",
			want: []pos{
				{1, 1},
				{2, 1},
				{3, 1},
				{0, 2},
				{1, 2},
				{0, 3},
				{1, 3},
				// reached EOF
				{1, 3},
				{1, 3},
				{1, 3},
			},
		},
		{
			name:  "Empty",
			input: "",
			want: []pos{
				// reached EOF
				{0, 1},
				{0, 1},
				{0, 1},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := NewBaseParser(
				source.NewStringSource(test.input),
			)

			for _, p := range test.want {
				assert.Equal(t, p.line, parser.LineNumber())
				assert.Equal(t, p.char, parser.CharNumber())
				parser.Next()
			}
		})
	}
}
