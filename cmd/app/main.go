package main

import (
	"fmt"

	"github.com/93mmm/proto-parser/internal/flags"
)

func main() {
	path := flags.DocPathOrDie()
	if err := RunParser(path); err != nil {
		fmt.Println(err)
	}
}
