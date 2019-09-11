package processor

type opcodeDecodeFunc func(byte, []byte) (instruction, error)

// NOP, HLT
func noParameterInstructionDecoder(
	opcode byte,
	parameters []byte,
) (instruction, error) {
	return instruction{opcode: opcode}, nil
}

// PRN
func oneRegisterInstructionDecoder(
	opcode byte,
	parameters []byte,
) (instruction, error) {
	return instruction{
		opcode: opcode,

		r1: parameters[0],
	}, nil
}

// LDM, STR, SWP, EQL, NQL
func twoRegisterInstructionDecoder(
	opcode byte,
	parameters []byte,
) (instruction, error) {
	return instruction{
		opcode: opcode,

		r1: parameters[0],
		r2: parameters[1],
	}, nil
}

// ADD. SUB, MUL, DIV
func threeRegisterInstructionDecoder(
	opcode byte,
	parameters []byte,
) (instruction, error) {
	return instruction{
		opcode: opcode,

		r1: parameters[0],
		r2: parameters[1],
		r3: parameters[2],
	}, nil
}

// LDI
func oneImmediateOneRegisterInstructionDecoder(
	opcode byte,
	parameters []byte,
) (instruction, error) {
	return instruction{
		opcode: opcode,

		r1: parameters[1],

		imm: parameters[0],
	}, nil
}

// JMP, JMC, JME
func oneImmediateInstructionDecoder(
	opcode byte,
	parameters []byte,
) (instruction, error) {
	return instruction{
		opcode: opcode,

		imm: parameters[0],
	}, nil
}

// unknown instruction
func unknownOpcodeDecoder(
	opcode byte,
	parameters []byte,
) (instruction, error) {
	return instruction{opcode: opcode}, nil
}
