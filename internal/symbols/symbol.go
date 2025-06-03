package symbols

import (
	"fmt"

	"github.com/93mmm/proto-parser/internal/token"
)

type Symbol struct {
	Name  string
	Type  string
	Line  int
	Start int
	End   int
}

func NewSymbol(name, kind string, line, start, end int) *Symbol {
	return &Symbol{
		Name:  name,
		Type:  kind,
		Line:  line,
		Start: start,
		End:   end,
	}
}

func (s *Symbol) String() string {
	actualType := s.Type
	if actualType == token.Rpc {
		actualType = "method"
	}

	return fmt.Sprintf(
		"%v %v %v:%v-%v",
		s.Name,
		actualType,
		s.Line,
		s.Start,
		s.End,
	)
}
