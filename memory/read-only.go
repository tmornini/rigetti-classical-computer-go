package memory

import (
	"fmt"
	"os"
)

// ReadOnly is memory that can be read from
type ReadOnly [256]byte

func (
	readOnly ReadOnly,
) Read(
	address Address,
	length int,
) ([]byte, error) {
	end := int(address) + length

	if end > 256 {
		fmt.Fprintf(
			os.Stderr,
			"out of bounds memory read on read-only memory: %x-%x\n",
			address,
			end,
		)
		return nil, ErrIllegalMemoryAccess
	}

	return readOnly[address:end], nil
}
