package main

import (
	"fmt"

	"github.com/93mmm/proto-parser/internal/flags"
)

func main() {
	docs := flags.DocPaths()
	if err := RunParser(docs); err != nil {
		fmt.Println(err)
	}
}
