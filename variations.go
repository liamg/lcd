package lcd

func New0802(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(8, 2, fontSize, registerSelect, enable, readWrite, data...)
}

func New1202(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(12, 2, fontSize, registerSelect, enable, readWrite, data...)
}

func New1601(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(16, 1, fontSize, registerSelect, enable, readWrite, data...)
}

func New1602(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(16, 2, fontSize, registerSelect, enable, readWrite, data...)
}

func New1604(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(16, 4, fontSize, registerSelect, enable, readWrite, data...)
}

func New2001(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(20, 1, fontSize, registerSelect, enable, readWrite, data...)
}

func New2002(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(20, 2, fontSize, registerSelect, enable, readWrite, data...)
}

func New2004(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(20, 4, fontSize, registerSelect, enable, readWrite, data...)
}

func New4002(fontSize FontSize, registerSelect, enable Pin, readWrite Pin, data ...Pin) (*LCD, error) {
	return New(40, 2, fontSize, registerSelect, enable, readWrite, data...)
}
