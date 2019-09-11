package processor

type instruction struct {
	opcode byte

	r1 byte
	r2 byte
	r3 byte

	imm byte
}

type instructionFunc func(*Processor, instruction) (programCounterAdvance int, err error)
