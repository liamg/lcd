package lcd1602

import (
	"fmt"
	"sync"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type LCD struct {
	displayMode byte
	pins        pins
}

type pins struct {
	registerSelect rpio.Pin
	readWrite      rpio.Pin
	enable         rpio.Pin
	data           []rpio.Pin
}

var openCount int
var openMu sync.Mutex

type instruction byte

const (
	insClearDisplay instruction = 1 << iota
	insReturnHome
	insEntryModeSet
	insSetDisplayMode
	insSetCursorDisplayShift
	insSetFunction
	insSetCGRAMAddress
	insSetDDRAMAddress
)

func Connect(registerSelect, readWrite, enable byte, data ...byte) (*LCD, error) {

	if len(data) != 4 && len(data) != 8 {
		return nil, fmt.Errorf("you must specify either 4 or 8 data pins")
	}

	openMu.Lock()
	defer openMu.Unlock()

	if openCount == 0 {
		if err := rpio.Open(); err != nil {
			return nil, err
		}
		openCount++
	}

	lcd := &LCD{
		displayMode: 0,
		pins: pins{
			registerSelect: rpio.Pin(registerSelect),
			readWrite:      rpio.Pin(readWrite),
			enable:         rpio.Pin(enable),
		},
	}

	for _, dataPin := range data {
		lcd.pins.data = append(lcd.pins.data, rpio.Pin(dataPin))
	}

	lcd.init()

	return lcd, nil
}

func (l *LCD) Close() error {
	openMu.Lock()
	defer openMu.Unlock()
	openCount--
	if openCount == 0 {
		return rpio.Close()
	}
	return nil
}

func (l *LCD) init() {
	l.pins.registerSelect.Low()
	l.pins.readWrite.Low()

	// -> 0x30
	l.pins.data[0].Low()
	l.pins.data[1].Low()
	l.pins.data[2].High()
	l.pins.data[3].High()

	time.Sleep(time.Millisecond * 5)

	// -> 0x30
	l.pins.data[0].Low()
	l.pins.data[1].Low()
	l.pins.data[2].High()
	l.pins.data[3].High()

	time.Sleep(time.Microsecond * 150)

	// -> 0x30
	l.pins.data[0].Low()
	l.pins.data[1].Low()
	l.pins.data[2].High()
	l.pins.data[3].High()

	// TODO: set config up

}

func (l *LCD) clearDisplay() {
	l.execInstruction(insSetDDRAMAddress, 0)
}

func (l *LCD) setDDRAMAddress(address byte) {
	l.execInstruction(insSetDDRAMAddress, address)
}

func (l *LCD) setCGRAMAddress(address byte) {
	l.execInstruction(insSetCGRAMAddress, address)
}

func (l *LCD) waitUntilFree() {
	l.pins.data[0].High()       // set db7 for bf
	l.pins.registerSelect.Low() // ins mode
	l.pins.readWrite.High()     // read
	for {
		if l.pins.data[0].Read() == rpio.Low {
			break
		}
		// keep firing the bf read instruction until the flag is unset
		l.pulseEnable()
	}
	//reset rw to default to write
	l.pins.readWrite.Low()
}

func (l *LCD) execInstruction(ins instruction, data byte) {
	l.waitUntilFree()
	l.pins.registerSelect.Low()
	data = byte(ins) | (data & (byte(ins) - 1))
	l.writeRaw(data)
}

func (l *LCD) writeData(data byte) {
	l.waitUntilFree()
	l.pins.registerSelect.High()
	l.writeRaw(data)
}

func (l *LCD) writeRaw(data byte) {
	if len(l.pins.data) == 8 {
		l.write8Bit(data)
		return
	}
	l.write4Bit(data)
}

func (l *LCD) write4Bit(data byte) {
	for i, pin := range l.pins.data {
		if data&(1<<(i+4)) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable()
	for i, pin := range l.pins.data {
		if data&(1<<i) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable()
}

func (l *LCD) write8Bit(data byte) {
	for i, pin := range l.pins.data {
		if data&(1<<i) > 0 {
			pin.High()
		} else {
			pin.Low()
		}
	}
	l.pulseEnable()
}

func (l *LCD) pulseEnable() {
	time.Sleep(time.Nanosecond * 500)
	l.pins.enable.High()
	time.Sleep(time.Nanosecond * 500)
	l.pins.enable.Low()
}

func (l *LCD) ShowCursor() {

}

func (l *LCD) setDisplayMode() {

}
