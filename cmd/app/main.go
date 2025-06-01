package main

import (
	"fmt"

	"github.com/93mmm/proto-parser/internal/flags"
)

func main() {
	args := flags.ParseArguments()
	fmt.Println(args)
}
