package processor

import "errors"

// ErrDivideByZero divide by zero
var ErrDivideByZero = errors.New("divide by zero")

// ErrUnknownOpcode unknown opcode
var ErrUnknownOpcode = errors.New("unknown opcode")

// ErrUnknownRegister unknown register
var ErrUnknownRegister = errors.New("unknown register")

// ErrUnknownFlag unknown flag
var ErrUnknownFlag = errors.New("unknown fkag")

// ErrHLTExecuted HLT instruction
var ErrHLTExecuted = errors.New("HLT instruction")

func errorIsNotContinuable(err error) bool {
	return err != nil && err != ErrDivideByZero
}
