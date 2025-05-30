package symbols

import "fmt"

type Symbol struct {
	name    string
	kind    string // Stands for type, but type is reserved
	line    int
	startID int
	endID   int
}

func (s *Symbol) String() string {
	return fmt.Sprintf(
		"%v %v %v:%v-%v",
		s.name,
		s.kind,
		s.line,
		s.startID,
		s.endID,
	)
}
