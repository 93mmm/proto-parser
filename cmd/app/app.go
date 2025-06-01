package main

import (
	"errors"
	"fmt"

	"github.com/93mmm/proto-parser/internal/parser"
	"github.com/93mmm/proto-parser/internal/source"
	"github.com/93mmm/proto-parser/internal/symbols"
)

func filterPrint(elements []*symbols.Symbol, ok func(*symbols.Symbol) bool) {
	for _, el := range elements {
		if ok(el) {
			fmt.Println(el)
		}
	}
}

func RunParser(docs []string) error {
	filter := func() func(*symbols.Symbol) bool {
		filterMap := map[string]struct{}{
			"import":  {},
			"service": {},
			"rpc":     {},
			"enum":    {},
			"message": {},
		}
		return func(s *symbols.Symbol) bool {
			_, ok := filterMap[s.Type()]
			return ok
		}
	}
	var errs []error

	for _, document := range docs {
		file, err := source.NewFileSource(document)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		defer file.Close()

		parsed, err := parser.NewParser(file).ParseDocument()
		if err != nil {
			fmt.Printf("An error occured while parsing %v: %v", document, err)
		}
		filterPrint(parsed, filter())
	}

	return errors.Join(errs...)
}
