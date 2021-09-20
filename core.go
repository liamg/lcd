package lcd

import (
	"time"
)

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
	l.pulseEnable(false, false)
	for i, pin := range l.pins.data {
		if data&(1<<i) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable(true, true)
}

func (l *LCD) write4BitHigh(data byte) {
	// send most significant 4 bits to most significant 4 pins
	for i, pin := range l.pins.data[len(l.pins.data)-4:] {
		if data&(1<<(i+4)) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable(true, false)
}

func (l *LCD) write8Bit(data byte) {
	for i, pin := range l.pins.data {
		if data&(1<<i) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable(true, true)
}

func (l *LCD) pulseEnable(withExecDelay bool, allowBF bool) {
	time.Sleep(time.Microsecond)
	l.pins.enable.High()
	time.Sleep(time.Microsecond)
	l.pins.enable.Low()
	if withExecDelay {
		l.wait(allowBF)
	}
}

func (l *LCD) wait(allowBF bool) {
	if l.pins.readWrite == nil || !allowBF {
		time.Sleep(timeExecutionDelay)
		return
	}

	l.pins.readWrite.High()
	defer l.pins.readWrite.Low()
	l.pins.setDataInput()
	defer l.pins.setDataOutput()

	for {
		l.pulseEnable(false, false)
		if !l.pins.data[len(l.pins.data)-1].Read() {
			break
		}
	}
}
