package processor

import (
	"fmt"
	"os"

	"github.com/tmornini/rigetti-computing/memory"
)

// Processor represents the Rigetti Classical Computer
type Processor struct {
	instructionSet instructionSet
	programMemory  *memory.ReadOnly
	mainMemory     *memory.ReadWrite

	programCounter memory.Address

	registers [4]byte // 0x00-0x03

	flags [6]bool // 0x04-0x05
}

// Boot create a new processor and make it process
func Boot(
	instructionSet instructionSet,
	programMemory *memory.ReadOnly,
	mainMemory *memory.ReadWrite,
) error {
	p := &Processor{
		instructionSet: instructionSet,
		programMemory:  programMemory,
		mainMemory:     mainMemory,
	}

	err := p.process()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(p)

	return err
}

func (p *Processor) process() error {
	for {
		instruction, err := p.decodeInstruction()
		if err != nil {
			p.flags[e] = true

			return err
		}

		programCounterAdvance, err := p.execute(instruction)
		if err != nil {
			if err == ErrHLTExecuted {
				return nil
			}

			p.flags[e] = true

			if errorIsNotContinuable(err) {
				return err
			}
		}

		p.programCounter += memory.Address(programCounterAdvance)
	}
}

func (p *Processor) execute(i instruction) (programCounterAdvance int, err error) {
	return p.instructionSet[i.opcode](p, i)
}

func (p *Processor) decodeInstruction() (instruction, error) {
	opcodeBytes, err := p.programMemory.Read(p.programCounter, 1)

	if err != nil {
		return instruction{}, err
	}

	opcode := opcodeBytes[0]

	parameterBytes, err := p.programMemory.Read(
		p.programCounter+1,
		opcodeParameterLengths[opcode],
	)
	if err != nil {
		return instruction{}, err
	}

	return opcodeDecodeFuncs[opcode](opcode, parameterBytes)
}

func (p Processor) registersAndFlagsAsString() string {
	return fmt.Sprintf(
		"PC:%x   X:%x   Y:%x   Z:%x   W:%x   C:%s   E:%s",
		[]byte{byte(p.programCounter)},
		[]byte{p.registers[x]},
		[]byte{p.registers[y]},
		[]byte{p.registers[z]},
		[]byte{p.registers[w]},
		string(fmt.Sprintf("%t", p.flags[c])[0]),
		string(fmt.Sprintf("%t", p.flags[e])[0]),
	)
}

func (p Processor) String() string {
	output := "Registers and Flags:\n"
	output += p.registersAndFlagsAsString() + "\n\n"

	output += "Program memory:\n"
	output += fmt.Sprintf("%x\n\n", *p.programMemory)

	output += "Main memory:\n"
	output += fmt.Sprintf("%x", *p.mainMemory)

	return output
}

var opcodeParameterLengths = [256]int{
	0, // NOP
	3, // ADD
	3, // SUB
	3, // MUL
	3, // DIV
	2, // LDM
	2, // LDI
	2, // STR
	2, // SWP
	2, // EQL
	2, // NQL
	1, // JMP
	1, // JMC
	1, // JME
	1, // PRN
	0, // HLT

	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var opcodeDecodeFuncs = [256]opcodeDecodeFunc{
	noParameterInstructionDecoder,             // NOP
	threeRegisterInstructionDecoder,           // ADD
	threeRegisterInstructionDecoder,           // SUB
	threeRegisterInstructionDecoder,           // MUL
	threeRegisterInstructionDecoder,           // DIV
	twoRegisterInstructionDecoder,             // LDM
	oneImmediateOneRegisterInstructionDecoder, // LDI
	twoRegisterInstructionDecoder,             // STR
	twoRegisterInstructionDecoder,             // SWP
	twoRegisterInstructionDecoder,             // EQL
	twoRegisterInstructionDecoder,             // NQL
	oneImmediateInstructionDecoder,            // JMP
	oneImmediateInstructionDecoder,            // JMC
	oneImmediateInstructionDecoder,            // JME
	oneRegisterInstructionDecoder,             // PRN
	noParameterInstructionDecoder,             // HLT

	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
	unknownOpcodeDecoder, unknownOpcodeDecoder, unknownOpcodeDecoder,
}
