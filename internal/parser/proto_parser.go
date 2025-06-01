package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/source"
)

type protoParser struct {
	base.BaseParser
}

func newProtoParser(src source.Source) *protoParser {
	p := &protoParser{
		BaseParser: *base.NewBaseParser(src),
	}
	return p
}
