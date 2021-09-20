package virtual

type DDRAM struct {
	data [80]byte
}

func (d *DDRAM) reset() {
	for i := 0; i < len(d.data); i++ {
		d.data[i] = 0x20
	}
}

func (d *DDRAM) isAddressOK(address byte) bool {
	return (address >= 0 && address <= 0x27) || (address >= 0x40 && address <= 0x67)
}

func (d *DDRAM) translateAddress(address byte) int {
	if address <= 0x27 {
		return int(address)
	}
	return int(address) - 0x18
}

func (d *DDRAM) write(address byte, data byte) {
	if !d.isAddressOK(address) {
		return
	}
	d.data[d.translateAddress(address)] = data
}

func (d *DDRAM) read(address byte) byte {
	if !d.isAddressOK(address) {
		return 0
	}
	return d.data[d.translateAddress(address)]
}
