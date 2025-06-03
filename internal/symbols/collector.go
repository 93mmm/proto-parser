package symbols

type Collector interface {
	Add(...*Symbol)
	All() []*Symbol
}

type collector struct {
	symbols []*Symbol
}

func NewCollector(capacity int) *collector {
	return &collector{
		symbols: make([]*Symbol, 0, capacity),
	}
}

func (c *collector) Add(s ...*Symbol) {
	c.symbols = append(c.symbols, s...)
}

func (c *collector) All() []*Symbol {
	return c.symbols
}
