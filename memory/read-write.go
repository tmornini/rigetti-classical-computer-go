package memory

import (
	"fmt"
	"os"
)

// ReadWrite is memory that can be read from and written to
type ReadWrite [256]byte

func (
	readWrite *ReadWrite,
) Read(
	address Address,
	length int,
) ([]byte, error) {
	end := int(address) + length

	if address < 0 || end > 256 {
		fmt.Fprintf(
			os.Stderr,
			"out of bounds read on read/write memory: %x-%x\n",
			address,
			int(end),
		)
		return nil, ErrIllegalMemoryAccess
	}

	return readWrite[address:end], nil
}

func (readWrite *ReadWrite) Write(address Address, bytes []byte) error {
	length := len(bytes)
	end := int(address) + length

	if address < 0 || end > 256 {
		fmt.Fprintf(
			os.Stderr,
			"out of bounds write on read/write memory: %x-%x\n",
			address,
			end,
		)
		return ErrIllegalMemoryAccess
	}

	for index, value := range bytes {
		(*readWrite)[address+Address(index)] = value
	}

	return nil
}
