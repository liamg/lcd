package lcd

import (
	"fmt"
	"time"
)

type LCD struct {
	columns     uint8
	lines       uint8
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
	if col >= l.columns {
		return fmt.Errorf("column number must be in the range 0-%d", l.columns-1)
	}
	if line >= l.lines {
		return fmt.Errorf("line number must be in the range 0-%d", l.lines-1)
	}
	l.setDDRAMAddress(l.addressForLine(line) + col)
	return nil
}

// WriteString writes string data to the current address
func (l *LCD) WriteString(data string) {
	l.pins.registerSelect.High()
	l.writeRaw([]byte(data)...)
}

// WriteLine writes a string to the given line of the display. Existing line content will be removed.
func (l *LCD) WriteLine(line uint8, text string) {
	l.execInstruction(insSetDDRAMAddress, l.addressForLine(line))
	text = fmt.Sprintf(fmt.Sprintf("%% -%ds", l.columns), text)
	for pos, c := range []byte(text) {
		if pos == int(l.columns) {
			break
		}
		l.Write(c)
	}
}
