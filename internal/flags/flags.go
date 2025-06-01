package flags

import (
	"fmt"
	"os"
)

var (
	errorMessage = "Error: no input arguments provided.\n"

	helpMessage = `Usage:
  protosym <input-file>

Example:
  protosym input.proto

Options:
  --help, -h     Show this help message

Please provide at least one input file to process.`
)

// I used to think about external library, but I think it is overhead for this tiny app
func DocPathOrDie() string {
	if len(os.Args) == 1 {
		fmt.Println(errorMessage + helpMessage)
		os.Exit(1)
	}

	path := os.Args[1]
	if path == "--help" || path == "-h" {
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	return path
}
