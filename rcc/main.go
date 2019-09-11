package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tmornini/rigetti-computing/memory"
	"github.com/tmornini/rigetti-computing/processor"
)

func main() {
	var err error
	var programReader io.ReadCloser

	switch len(os.Args) {
	case 1:
		programReader = os.Stdin
	case 2:
		programReader, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(
			os.Stderr,
			"usage: "+os.Args[0]+"[256 byte binary file]",
		)
		os.Exit(2)
	}

	programMemory, err := memory.NewProgramFrom(programReader)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	instructionSet := processor.NormalInstructionSet

	if os.Getenv("DEBUG") != "" {
		instructionSet = processor.DebugInstructionSet

		if os.Getenv("NONOP") != "" {
			instructionSet[0] = processor.NormalInstructionSet[0]
		}
	}

	err = processor.Boot(
		instructionSet,
		programMemory,
		&memory.ReadWrite{},
	)
	if err != nil {
		os.Exit(4)
	}
}
