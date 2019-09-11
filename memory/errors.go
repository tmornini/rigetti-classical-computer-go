package memory

import "errors"

// ErrIllegalMemoryAccess illegal memory access
var ErrIllegalMemoryAccess = errors.New("illegal memory access")

// ErrInvalidProgramLength invalid program length
var ErrInvalidProgramLength = errors.New("invalid program length")
