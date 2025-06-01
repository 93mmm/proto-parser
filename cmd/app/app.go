package main

import (
	"fmt"

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
	file, err := source.NewFileSource(document)
	if err != nil {
		return err
	}
	defer file.Close()

	parsed, err := parser.NewParser(file).ParseDocument()
	if err != nil {
		return err
	}

	for _, element := range parsed {
		filterPrint(element)
	}
	return nil
}
