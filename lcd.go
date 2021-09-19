package lcd1602

import (
	"fmt"
	"time"
)

type LCD struct {
	displayMode displayMode
	entryMode   entryMode
	pins        pins
}

// Clear removes all content by overwriting all characters with spaces and moving the cursor to 0x00
func (l *LCD) Clear() {
	l.execInstruction(insClearDisplay, 0)
	time.Sleep(timeClearDelay)
}

// Write writes data to the current address
func (l *LCD) Write(data ...byte) {
	l.pins.registerSelect.High()
	l.writeRaw(data...)
}

// MoveTo relocates the cursor to the given position
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

// WriteString writes string data to the current address
func (l *LCD) WriteString(data string) {
	l.pins.registerSelect.High()
	l.writeRaw([]byte(data)...)
}

// WriteTopLine writes a string to the top line of the display. Existing content will be removed.
func (l *LCD) WriteTopLine(text string) {
	l.execInstruction(insSetDDRAMAddress, addrLine1)
	text = fmt.Sprintf("% -16s", text)
	for pos, c := range []byte(text) {
		if pos == 16 {
			break
		}
		l.Write(c)
	}
}

// WriteBottomLine writes a string to the top line of the display. Existing content will be removed.
func (l *LCD) WriteBottomLine(text string) {
	l.execInstruction(insSetDDRAMAddress, addrLine2)
	text = fmt.Sprintf("% -16s", text)
	for pos, c := range []byte(text) {
		if pos == 16 {
			break
		}
		l.Write(c)
	}
}
