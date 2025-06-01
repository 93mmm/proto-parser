package parser

import (
	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/source"
)

type protoParser struct {
	*base.BaseParser
}

func NewProtoParser(bp *base.BaseParser) *protoParser {
	return &protoParser{
		BaseParser: bp,
	}

}

func newTestProtoParser(input string) *protoParser {
	src := source.NewStringSource(input)
	bp := base.NewBaseParser(src)
	return NewProtoParser(bp)
}
