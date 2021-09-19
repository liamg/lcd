package lcd

func New0802(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(8, 2, registerSelect, enable, data...)
}

func New1202(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(12, 2, registerSelect, enable, data...)
}

func New1601(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(16, 1, registerSelect, enable, data...)
}

func New1602(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(16, 2, registerSelect, enable, data...)
}

func New1604(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(16, 4, registerSelect, enable, data...)
}

func New2001(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(20, 1, registerSelect, enable, data...)
}

func New2002(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(20, 2, registerSelect, enable, data...)
}

func New2004(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(20, 4, registerSelect, enable, data...)
}

func New4002(registerSelect, enable Pin, data ...Pin) (*LCD, error) {
	return New(40, 2, registerSelect, enable, data...)
}
