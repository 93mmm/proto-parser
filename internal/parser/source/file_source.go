package source

import (
	"bufio"
	"os"
)

type fileSource struct {
	src  *bufio.Reader
	file *os.File
	pos  int
}

func NewFileSource(filename string) (*fileSource, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	return &fileSource{
		src:  reader,
		file: file,
	}, nil
}

func (s *fileSource) Next() (rune, error) {
	r, _, err := s.src.ReadRune()
	if err != nil {
		return EOF, err
	}
	return r, nil
}

func (s *fileSource) Close() error {
	return s.file.Close()
}
