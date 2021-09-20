package virtual

import (
	"time"
)

const (
	insClearDisplay byte = 1 << iota
	insReturnHome
	insEntryModeSet
	insSetDisplayMode
	insSetCursorDisplayShift
	insSetFunction
	insSetCGRAMAddress
	insSetDDRAMAddress
)

func (l *LCD) setBusy(b bool) {
	l.busy = b
}

func (l *LCD) processInstruction(instruction byte) {

	if l.initialised && l.rwPin.state {
		l.insReadBusyFlag()
		return
	}

	l.busyMu.Lock()
	defer l.busyMu.Unlock()
	l.setBusy(true)
	defer l.setBusy(false)

	time.Sleep(time.Microsecond * 40)

	if instruction&insSetDDRAMAddress > 0 {
		l.insSetDDRAMAddress(instruction)
		return
	}
	if instruction&insSetCGRAMAddress > 0 {
		l.insSetCGRAMAddress(instruction)
		return
	}
	if instruction&insSetFunction > 0 {
		l.insSetFunc(instruction)
		return
	}
	if instruction&insSetCursorDisplayShift > 0 {
		l.insShift(instruction)
		return
	}
	if instruction&insSetDisplayMode > 0 {
		l.insSetDisplayMode(instruction)
		return
	}
	if instruction&insEntryModeSet > 0 {
		l.insSetEntryMode(instruction)
		return
	}
	if instruction&insReturnHome > 0 {
		l.insReturnHome()
		return
	}
	if instruction&insClearDisplay > 0 {
		l.insClearDisplay()
		return
	}
}

func (l *LCD) insClearDisplay() {
	l.ddram.reset()
}

func (l *LCD) insReturnHome() {
	l.currentAddress = address{
		location: 0,
		kind:     addressKindDDRAM,
	}
}

func (l *LCD) insSetEntryMode(data byte) {
	l.config.increment = data&0b10 > 0
	l.config.shift = data&0b1 > 0
}

func (l *LCD) insSetDisplayMode(data byte) {
	l.config.on = data&0b100 > 0
	l.config.cursor = data&0b10 > 0
	l.config.blink = data&0b1 > 0
}

func (l *LCD) insShift(data byte) {
	panic("not implemented")
}

func (l *LCD) insSetCGRAMAddress(data byte) {
	l.currentAddress = address{
		kind:     addressKindCGRAM,
		location: data & 0b111111,
	}
}

func (l *LCD) insSetDDRAMAddress(data byte) {
	l.currentAddress = address{
		kind:     addressKindDDRAM,
		location: data & 0b1111111,
	}
}

func (l *LCD) insReadBusyFlag() {
	// set address
	l.writePins(l.currentAddress.location)
	// set busy pin
	if l.busy {
		l.dataPins[len(l.dataPins)-1].High()
	} else {
		l.dataPins[len(l.dataPins)-1].Low()
	}
}

func (l *LCD) insSetFunc(data byte) {
	switch l.initStage {
	case 0:
		l.initStage++
		l.lastInitStage = time.Now()
	case 1:
		if time.Since(l.lastInitStage) >= 41*(time.Millisecond/10) {
			l.initStage++
			l.lastInitStage = time.Now()
		}
	case 2:
		if time.Since(l.lastInitStage) >= 100*time.Microsecond {
			l.initStage++
			l.lastInitStage = time.Now()
		}
	case 3:
		l.initStage = 0
		l.initialised = true
	}

	if !l.initialised {
		return
	}

	// process "real" function set
	if data&0b10000 > 0 { // data length
		l.config.busSize = 8
	} else {
		l.config.busSize = 4
	}

	if data&0b1000 > 0 { // N number of lines
		l.config.lines = 2
	} else {
		l.config.lines = 1
	}

	if data&0b100 > 0 { // F font 5x11/5x8
		l.config.fontHeight = 11
	} else {
		l.config.fontHeight = 8
	}
}
