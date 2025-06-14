package source

import (
	"io"

	"github.com/93mmm/proto-parser/internal/parser/constants"
)

type stringSource struct {
	src []rune
	pos int
}

func NewStringSource(s string) *stringSource {
	return &stringSource{
		src: []rune(s),
	}
}

func (s *stringSource) Next() (rune, error) {
	if s.pos >= len(s.src) {
		return constants.EOF, io.EOF
	}
	r := s.src[s.pos]
	s.pos++
	return r, nil
}
