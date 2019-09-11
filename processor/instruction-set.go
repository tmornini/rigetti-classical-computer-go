package processor

import (
	"fmt"

	"github.com/tmornini/rigetti-computing/memory"
)

type instructionSet [256]instructionFunc

func nop(p *Processor, i instruction) (programCounterAdvance int, err error) {
	return 1, nil
}

func add(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) || unknownRegister(i.r3) {
		return 4, ErrUnknownRegister
	}

	p.registers[i.r3] = p.registers[i.r1] + p.registers[i.r2]

	return 4, nil
}

func sub(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) || unknownRegister(i.r3) {
		return 4, ErrUnknownRegister
	}

	p.registers[i.r3] = p.registers[i.r1] - p.registers[i.r2]

	return 4, nil
}

func mul(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) || unknownRegister(i.r3) {
		return 4, ErrUnknownRegister
	}

	p.registers[i.r3] = p.registers[i.r1] * p.registers[i.r2]

	return 4, nil
}

func div(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) || unknownRegister(i.r3) {
		return 4, ErrUnknownRegister
	}

	if p.registers[i.r2] == 0 {
		return 4, ErrDivideByZero
	}

	p.registers[i.r3] = p.registers[i.r1] / p.registers[i.r2]
	if err != nil {
		return 4, err
	}

	return 4, nil
}

func ldm(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) {
		return 3, ErrUnknownRegister
	}

	address := memory.Address(p.registers[i.r1])

	bytes, err := p.mainMemory.Read(address, 1)
	if err != nil {
		return 3, err
	}

	p.registers[i.r2] = bytes[0]

	return 3, nil
}

func ldi(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) {
		return 3, ErrUnknownRegister
	}

	p.registers[i.r1] = i.imm

	return 3, nil
}

func str(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) {
		return 3, ErrUnknownRegister
	}

	r2Address := memory.Address(p.registers[i.r2])

	err = p.mainMemory.Write(r2Address, []byte{p.registers[i.r1]})
	if err != nil {
		return 3, err
	}

	return 3, nil
}

func swp(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) {
		return 3, ErrUnknownRegister
	}

	p.registers[i.r1], p.registers[i.r2] = p.registers[i.r2], p.registers[i.r1]

	return 3, nil
}

func eql(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) {
		return 3, ErrUnknownRegister
	}

	p.flags[c] = p.registers[i.r1] == p.registers[i.r2]

	return 3, nil
}

func nql(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) || unknownRegister(i.r2) {
		return 3, ErrUnknownRegister
	}

	p.flags[c] = p.registers[i.r1] != p.registers[i.r2]

	return 3, nil
}

func jmp(p *Processor, i instruction) (programCounterAdvance int, err error) {
	p.programCounter = memory.Address(i.imm)

	return 0, nil
}

func jmc(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if p.flags[c] {
		p.programCounter = memory.Address(i.imm)
		p.flags[c] = false
		return 0, nil
	}

	return 2, nil
}

func jme(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if p.flags[e] {
		p.programCounter = memory.Address(i.imm)
		p.flags[e] = false
		return 0, nil
	}

	return 2, nil
}

func prn(p *Processor, i instruction) (programCounterAdvance int, err error) {
	if unknownRegister(i.r1) {
		return 2, ErrUnknownRegister
	}

	fmt.Print(
		string(
			p.registers[i.r1],
		),
	)

	return 2, nil
}

func hlt(p *Processor, i instruction) (programCounterAdvance int, err error) {
	return 1, ErrHLTExecuted
}

func unknown(p *Processor, i instruction) (programCounterAdvance int, err error) {
	return 0, ErrUnknownOpcode
}

// NormalInstructionSet is the non-debugging instruction set
var NormalInstructionSet = instructionSet{
	nop, add, sub, mul, div, ldm, ldi, str,
	swp, eql, nql, jmp, jmc, jme, prn, hlt,

	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
	unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown,
}
