package lcd

import (
	"fmt"
	"time"
)

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
	l.execInstruction(insSetDisplayMode, l.displayMode.byte())

	// Clear
	l.Clear()

	// Entry mode
	l.SetCursorIncrement(true)

	// Turn on!
	l.On()
}

// New opens communication with the LCD for further instruction
func New(columns, lines uint8, registerSelect, enable Pin, data ...Pin) (*LCD, error) {

	if len(data) != 4 && len(data) != 8 {
		return nil, fmt.Errorf("you must specify either 4 or 8 data pins")
	}

	lcd := &LCD{
		columns: columns,
		lines:   lines,
		pins: pins{
			registerSelect: registerSelect,
			enable:         enable,
			data:           data,
		},
	}

	lcd.init()

	return lcd, nil
}
