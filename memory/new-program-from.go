package memory

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// NewProgramFrom creates a new read-only memory from a reader
func NewProgramFrom(programReader io.Reader) (*ReadOnly, error) {
	programBytes, err := ioutil.ReadAll(programReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading code: %s\n", err)
		return nil, err
	}

	length := len(programBytes)

	if length < 1 || length > 256 {
		fmt.Fprintf(
			os.Stderr,
			"invalid program length, must be 1-256 bytes, is: %d\n",
			length,
		)

		return nil, ErrInvalidProgramLength
	}

	programMemory := &ReadOnly{}

	copy(programMemory[:], programBytes[0:256])

	return programMemory, nil
}
