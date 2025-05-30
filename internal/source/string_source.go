package source

import "io"

type stringSource struct {
	src []rune
	pos int
}

var _ Source = (*stringSource)(nil)

func NewStringSource(s string) *stringSource {
	return &stringSource{
		src: []rune(s),
	}
}

func (s *stringSource) Next() (rune, error) {
	if s.pos >= len(s.src) {
		return EOF, io.EOF
	}
	r := s.src[s.pos]
	s.pos++
	return r, nil
}

