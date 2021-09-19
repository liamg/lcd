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

const (
	addrLine1 = 0x00
	addrLine2 = 0x40
)

const (
	timeExecutionDelay = time.Microsecond * 80
	timeClearDelay     = time.Millisecond * 10
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

func (l *LCD) init4Bit() {
	l.write4BitHigh(0b00100000) // 4-bit mode - N/F flags ignored

	// from here we can send high then low nibble
	l.write4Bit(0b00101000) // 4 bit, 2 lines, 5x8
}

func (l *LCD) init8Bit() {
	l.write8Bit(0b00111000) // 8-bit mode, 2 lines, 5x8
}

func (l *LCD) init() {

	l.pins.init()

	l.pins.registerSelect.Low() // instruction
	l.pins.readWrite.Low()      // write

	l.write4BitHigh(0b00110000)
	time.Sleep(time.Millisecond * 10)

	l.write4BitHigh(0b00110000)
	time.Sleep(time.Microsecond * 200)

	l.write4BitHigh(0b00110000)
	time.Sleep(time.Millisecond * 10)

	if len(l.pins.data) == 4 {
		l.init4Bit()
	} else {
		l.init8Bit()
	}

	// Display OFF
	l.execInstruction(insSetDisplayMode, 0x0)

	// Clear
	l.Clear()

	// Entry mode
	l.execInstruction(insEntryModeSet, 0b010)

	// Turn on!
	l.execInstruction(insSetDisplayMode, 0b100)
}

func (l *LCD) Clear() {
	l.execInstruction(insClearDisplay, 0)
	time.Sleep(timeClearDelay)
}

func (l *LCD) MoveTo(line uint8, col uint8) error {
	if col >= 16 {
		return fmt.Errorf("column number must be in the range 0-15")
	}
	if line > 1 {
		return fmt.Errorf("line number must be in the range 0-1")
	}
	var address byte
	if line == 1 {
		address = addrLine2
	}
	l.setDDRAMAddress(address + col)
	return nil
}

func (l *LCD) setDDRAMAddress(address byte) {
	l.execInstruction(insSetDDRAMAddress, address)
}

func (l *LCD) setCGRAMAddress(address byte) {
	l.execInstruction(insSetCGRAMAddress, address)
}

func (l *LCD) waitUntilFree() {

	// TODO
	time.Sleep(timeExecutionDelay)
	return

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
		l.pulseEnable(true)
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
	l.pins.registerSelect.Low()
	data = byte(ins) | data
	l.writeRaw(data)
}

// Write writes data to the current address
func (l *LCD) Write(data ...byte) {
	l.pins.registerSelect.High()
	l.writeRaw(data...)
}

// WriteString writes string data to the current address
func (l *LCD) WriteString(data string) {
	l.pins.registerSelect.High()
	l.writeRaw([]byte(data)...)
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
		l.waitUntilFree()
	}
}

func (l *LCD) WriteTopLine(text string) {
	l.execInstruction(insSetDDRAMAddress, addrLine1)
	for pos, c := range []byte(text) {
		if pos == 16 {
			break
		}
		l.Write(c)
	}
}

func (l *LCD) WriteBottomLine(text string) {
	l.execInstruction(insSetDDRAMAddress, addrLine2)
	for pos, c := range []byte(text) {
		if pos == 16 {
			break
		}
		l.Write(c)
	}
}
