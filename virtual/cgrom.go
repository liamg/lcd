package virtual

type CGROM interface {
	read(address uint8) (character []byte)
}

type CGROM5x8 struct{}
type CGROM5x11 struct{}

var chars5x8 = map[uint8][8]byte{
	0x61: {
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
	},
}

var chars5x11 = map[uint8][11]byte{
	0x61: {
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
		0b00000,
	},
}

func (c *CGROM5x8) read(address uint8) []byte {
	data, ok := chars5x8[address]
	if !ok {
		return make([]byte, 8)
	}
	return data[:]
}

func (c *CGROM5x11) read(address uint8) []byte {
	data, ok := chars5x11[address]
	if !ok {
		return make([]byte, 11)
	}
	return data[:]
}
