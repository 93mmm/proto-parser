package flags

import (
	"fmt"
	"os"
)

var helpMessage = 
`Error: no input arguments provided.

Usage:
  protosym <input-files>...

Example:
  protosym input.proto file.proto

Options:
  --help, -h     Show this help message

Please provide at least one input file to process.`

func ParseArguments() []string {
	if len(os.Args) == 1 {
		fmt.Println(helpMessage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "--help", "-h":
		fmt.Println(helpMessage)
		os.Exit(1)
	}
	return os.Args[1:]
}
