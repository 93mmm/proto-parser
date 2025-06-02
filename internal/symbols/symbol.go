package symbols

import "fmt"

type Symbol struct {
	name  string
	kind  string // Stands for type, but type is reserved
	line  int
	start int
	end   int
}

func NewSymbol(name, kind string, line, start, end int) *Symbol {
	return &Symbol{
		name:  name,
		kind:  kind,
		line:  line,
		start: start,
		end:   end,
	}
}

func (s *Symbol) String() string {
	return fmt.Sprintf(
		"%v %v %v:%v-%v",
		s.name,
		s.kind,
		s.line,
		s.start,
		s.end,
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

func (s *Symbol) SetStart(start int) *Symbol {
	s.start = start
	return s
}

func (s *Symbol) SetEnd(end int) *Symbol {
	s.end = end
	return s
}

func (s *Symbol) Name() string { return s.name }
func (s *Symbol) Type() string { return s.kind }
