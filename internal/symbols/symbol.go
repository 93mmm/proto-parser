package symbols

import (
	"fmt"

	"github.com/93mmm/proto-parser/internal/token"
)

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
	actualType := s.kind
	if actualType == token.Rpc {
		actualType = "method"
	}

	return fmt.Sprintf(
		"%v %v %v:%v-%v",
		s.name,
		actualType,
		s.line,
		s.start,
		s.end,
	)
}

func (s *Symbol) Name() string { return s.name }
func (s *Symbol) Type() string { return s.kind }
