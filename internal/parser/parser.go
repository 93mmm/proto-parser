package parser

import (
	base "github.com/93mmm/protoc-parser/internal/baseparser"
	"github.com/93mmm/protoc-parser/internal/source"
)

type ProtocParser struct {
	base.BaseParser
	line int
	char int
}

func NewProtocParser(src source.Source) *ProtocParser {
	return &ProtocParser{
		BaseParser: *base.NewBaseParser(src),
	}
}
