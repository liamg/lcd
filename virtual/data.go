package virtual

func (l *LCD) processData(data byte) {
	l.busyMu.Lock()
	defer l.busyMu.Unlock()
	l.setBusy(true)
	defer l.setBusy(false)

	if l.rwPin.Read() {
		l.readData(data)
		return
	}

	l.writeData(data)
}

func (l *LCD) writeData(data byte) {
	switch l.currentAddress.kind {
	case addressKindDDRAM:
		l.ddram.write(l.currentAddress.location, data)
	case addressKindCGRAM:
		l.cgram.write(l.currentAddress.location, data)
	default:
		panic("invalid address kind")
	}
	l.currentAddress.location++
}

func (l *LCD) readData(data byte) {
	switch l.currentAddress.kind {
	case addressKindDDRAM:
		l.writePins(
			l.ddram.read(l.currentAddress.location),
		)
	case addressKindCGRAM:
		l.writePins(
			l.cgram.read(l.currentAddress.location),
		)
	default:
		panic("invalid address kind")
	}
}

func (l *LCD) writePins(data byte) {
	for i := 0; i < 8; i++ {
		if data&1<<i > 0 {
			l.dataPins[i].High()
		} else {
			l.dataPins[i].Low()
		}
	}
}

func (l *LCD) readPins() byte {
	var output byte
	for i := 0; i < 8; i++ {
		if l.dataPins[i].state {
			output |= (1 << i)
		}
	}
	return output
}
