package processor

const ( // iota is reset to 0
	x = iota
	y = iota
	z = iota
	w = iota

	c = iota
	e = iota
)

func unknownRegister(register byte) bool {
	if register > 3 {
		return true
	}

	return false
}
