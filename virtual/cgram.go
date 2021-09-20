package virtual

type CGRAM struct {
	data [64]byte // enough room for 8 * 5x8
}

func (c *CGRAM) write(address byte, data byte) {
	if int(address) >= len(c.data) {
		return
	}
	c.data[address] = data
}

func (c *CGRAM) read(address byte) byte {
	if int(address) >= len(c.data) {
		return 0
	}
	return c.data[address]
}
