package main

import (
	"fmt"

	base "github.com/93mmm/proto-parser/internal/baseparser"
	"github.com/93mmm/proto-parser/internal/lexer"
	"github.com/93mmm/proto-parser/internal/parser"
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
)

var filterMap = map[string]struct{}{
	"import":  {},
	"service": {},
	"rpc":     {},
	"enum":    {},
	"message": {},
}

func filterPrint(el *symbols.Symbol) {
	if _, ok := filterMap[el.Type()]; ok {
		fmt.Println(el)
	}
}

func RunParser(document string) error {
	src, err := source.NewFileSource(document)
	if err != nil {
		return err
	}
	defer src.Close()
	bp := base.NewBaseParser(src)
	l := lexer.NewLexer(bp)
	pp := parser.NewProtoParser(l)
	parsed, err := parser.NewParser(pp).ParseDocument()
	if err != nil {
		return err
	}

	for _, element := range parsed {
		filterPrint(element)
	}
	return nil
}
