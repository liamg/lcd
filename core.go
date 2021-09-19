package lcd

import "time"

const (
	timeExecutionDelay = time.Microsecond * 80
	timeClearDelay     = time.Millisecond * 10
)

func (l *LCD) setDDRAMAddress(address byte) {
	l.execInstruction(insSetDDRAMAddress, address)
}

func (l *LCD) setCGRAMAddress(address byte) {
	l.execInstruction(insSetCGRAMAddress, address)
}

func (l *LCD) execInstruction(ins instruction, data byte) {
	l.pins.registerSelect.Low()
	data = byte(ins) | data
	l.writeRaw(data)
}

func (l *LCD) writeRaw(data ...byte) {
	for _, b := range data {
		if len(l.pins.data) == 8 {
			l.write8Bit(b)
			return
		}
		l.write4Bit(b)
	}
}

func (l *LCD) write4Bit(data byte) {
	for i, pin := range l.pins.data {
		if data&(1<<(i+4)) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable(false)
	for i, pin := range l.pins.data {
		if data&(1<<i) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable(true)
}

func (l *LCD) write4BitHigh(data byte) {
	for i, pin := range l.pins.data {
		if data&(1<<(i+4)) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable(true)
}

func (l *LCD) write8Bit(data byte) {
	for i, pin := range l.pins.data {
		if data&(1<<i) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable(true)
}

func (l *LCD) pulseEnable(withExecDelay bool) {
	time.Sleep(time.Microsecond)
	l.pins.enable.High()
	time.Sleep(time.Microsecond)
	l.pins.enable.Low()
	if withExecDelay {
		time.Sleep(timeExecutionDelay)
	}
}
