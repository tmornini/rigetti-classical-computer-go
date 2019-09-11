package processor

import (
	"fmt"
	"os"
)

func nopDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.noParameterInstructionString(),
	)

	return nop(p, i)
}

func addDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.threeRegisterInstructionString(),
	)

	return add(p, i)
}

func subDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.threeRegisterInstructionString(),
	)

	return sub(p, i)
}

func mulDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.threeRegisterInstructionString(),
	)

	return mul(p, i)
}

func divDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.threeRegisterInstructionString(),
	)

	return div(p, i)
}

func ldmDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.twoRegisterInstructionString(),
	)

	return ldm(p, i)
}

func ldiDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.oneImmediateOneRegisterInstructionString(),
	)

	return ldi(p, i)
}

func strDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.twoRegisterInstructionString(),
	)

	return str(p, i)
}

func swpDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.twoRegisterInstructionString(),
	)

	return swp(p, i)
}

func eqlDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.twoRegisterInstructionString(),
	)

	return eql(p, i)
}

func nqlDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.twoRegisterInstructionString(),
	)

	return nql(p, i)
}

func jmpDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.oneImmediateInstructionString(),
	)

	return jmp(p, i)
}

func jmcDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.oneImmediateInstructionString(),
	)

	return jmc(p, i)
}

func jmeDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.oneImmediateInstructionString(),
	)

	return jme(p, i)
}

func prnDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.oneRegisterInstructionString(),
	)

	return prn(p, i)
}

func hltDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+
			"   |   "+
			i.noParameterInstructionString(),
	)

	return hlt(p, i)
}

func unknownDebug(p *Processor, i instruction) (programCounterAdvance int, err error) {
	fmt.Fprintln(
		os.Stderr,
		p.registersAndFlagsAsString()+"   |   ???",
	)

	return unknown(p, i)
}

func (i instruction) Name() string {
	return opcodeNames[i.opcode]
}

func (i instruction) registerName(r byte) string {
	return registerNames[r]
}

func (i instruction) noParameterInstructionString() string {
	return i.Name()
}

func (i instruction) oneRegisterInstructionString() string {
	return fmt.Sprintf(
		"%s %s",
		i.Name(),
		i.registerName(i.r1),
	)
}

func (i instruction) twoRegisterInstructionString() string {
	return fmt.Sprintf(
		"%s %s %s",
		i.Name(),
		i.registerName(i.r1),
		i.registerName(i.r2),
	)
}

func (i instruction) threeRegisterInstructionString() string {
	return fmt.Sprintf(
		"%s %s %s %s",
		i.Name(),
		i.registerName(i.r1),
		i.registerName(i.r2),
		i.registerName(i.r3),
	)
}

func (i instruction) oneImmediateOneRegisterInstructionString() string {
	return fmt.Sprintf(
		"%s %x %s",
		i.Name(),
		i.imm,
		i.registerName(i.r1),
	)
}

func (i instruction) oneImmediateInstructionString() string {
	return fmt.Sprintf("%s %x", i.Name(), i.imm)
}

var opcodeNames = [256]string{
	"NOP", "ADD", "SUB", "MUL", "DIV", "LDM", "LDI", "STR",
	"SWP", "EQL", "NQL", "JMP", "JMC", "JME", "PRN", "HLT",

	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
	"???", "???", "???", "???", "???", "???", "???", "???",
}

var registerNames = [256]string{
	"X", "Y", "Z", "W",

	"?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
	"?", "?", "?", "?", "?", "?", "?", "?",
}

// DebugInstructionSet wraps InstructionSetNormal to provide debug output
var DebugInstructionSet = instructionSet{
	nopDebug, addDebug, subDebug, mulDebug,
	divDebug, ldmDebug, ldiDebug, strDebug,
	swpDebug, eqlDebug, nqlDebug, jmpDebug,
	jmcDebug, jmeDebug, prnDebug, hltDebug,

	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
	unknownDebug, unknownDebug, unknownDebug, unknownDebug,
}
