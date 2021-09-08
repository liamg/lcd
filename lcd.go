package lcd1602

import (
	"fmt"
	"sync"

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

	lcd := LCD{
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

	return &lcd, nil
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

func (l *LCD) clearDisplay() {
	l.execInstruction(insSetDDRAMAddress, 0)
}

func (l *LCD) setDDRAMAddress(address byte) {
	l.execInstruction(insSetDDRAMAddress, address&0x7f)
}

func (l *LCD) setCGRAMAddress(address byte) {
	l.execInstruction(insSetCGRAMAddress, address&0x3f)
}

func (l *LCD) execInstruction(ins instruction, data byte) {
	l.pins.registerSelect.Low()

	// TODO set data pins

	_ = byte(ins) | (data & (byte(ins) - 1))
}

func (l *LCD) ShowCursor() {
}

func (l *LCD) setDisplayMode() {

}

func (l *LCD) writeData(data byte) {

}
