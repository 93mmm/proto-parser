package symbols

import "fmt"

type Symbol struct {
	name      string
	kind      string // Stands for type, but type is reserved
	line      int
	startChar int
	endChar   int
}

func NewSymbol(name string, kind string, line int, startChar int, endChar int) *Symbol {
	return &Symbol{
		name:      name,
		kind:      kind,
		line:      line,
		startChar: startChar,
		endChar:   endChar,
	}
}

func (s *Symbol) String() string {
	return fmt.Sprintf(
		"%v %v %v:%v-%v",
		s.name,
		s.kind,
		s.line,
		s.startChar,
		s.endChar,
	)
}

func (s *Symbol) SetName(name string) *Symbol { 
	s.name = name
	return s
}

func (s *Symbol) SetType(kind string) *Symbol {
	s.kind = kind
	return s
}

func (s *Symbol) SetLine(line int) *Symbol {
	s.line = line
	return s
}

func (s *Symbol) SetStartChar(startChar int) *Symbol {
	s.startChar = startChar
	return s
}

func (s *Symbol) SetEndChar(endChar int) *Symbol {
	s.endChar = endChar
	return s
}

func (s *Symbol) Name() string { return s.name }
func (s *Symbol) Type() string { return s.kind }
