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

func (p *pins) init() {
	p.registerSelect.Output()
	p.readWrite.Output()
	p.enable.Output()
	for _, data := range p.data {
		data.Output()
	}
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

	l.pins.init()

	l.pins.registerSelect.Low()
	l.pins.readWrite.Low()

	l.execInstruction(0x00, 0b00110000)

	time.Sleep(time.Millisecond * 10)

	l.execInstruction(0x00, 0b00110000)

	time.Sleep(time.Microsecond * 200)

	l.execInstruction(0x00, 0b00110000)

	time.Sleep(time.Microsecond * 200)

	// Function Set
	var fsParam byte
	if len(l.pins.data) == 8 {
		fsParam |= 0b10000
	}
	// 2 lines
	fsParam |= 0b1000
	l.execInstruction(insSetFunction, fsParam)

	// Display OFF
	l.execInstruction(insSetDisplayMode, 0x0)

	// Clear
	l.execInstruction(insClearDisplay, 0x0)

	// Entry mode
	l.execInstruction(insEntryModeSet, 0b010)

	// Turn on!
	l.execInstruction(insSetDisplayMode, 0b100)
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

	for _, pin := range l.pins.data {
		pin.Low()
	}

	db7 := l.pins.data[len(l.pins.data)-1]

	l.pins.registerSelect.Low() // ins mode
	l.pins.readWrite.High()     // read

	for {
		// keep firing the bf read instruction until the flag is unset
		db7.Output()
		db7.High()
		l.pulseEnable()
		db7.Input()
		if db7.Read() == rpio.Low {
			break
		}
	}

	db7.Output()
	//reset rw to default to write
	l.pins.readWrite.Low()
}

func (l *LCD) execInstruction(ins instruction, data byte) {
	l.waitUntilFree()
	l.pins.registerSelect.Low()
	data = byte(ins) | data
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

func (l *LCD) WriteLines(lines ...string) {
	l.writeData(0x65)
}
