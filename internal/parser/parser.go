package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/source"
)

type ProtoParser struct {
	base.BaseParser
	line int
	char int
}

func NewProtoParser(src source.Source) *ProtoParser {
	return &ProtoParser{
		BaseParser: *base.NewBaseParser(src),
	}
}
