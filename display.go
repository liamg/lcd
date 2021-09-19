package lcd1602

type displayMode struct {
	On     bool
	Cursor bool
	Blink  bool
}

func (d displayMode) byte() byte {
	var output byte
	if d.On {
		output |= 0b100
	}
	if d.Cursor {
		output |= 0b10
	}
	if d.Blink {
		output |= 0b1
	}
	return output
}

// On enables the LCD display
func (l *LCD) On() {
	l.displayMode.On = true
	l.execInstruction(insSetDisplayMode, l.displayMode.byte())
}

// Off disables the LCD display (content will be preserved)
func (l *LCD) Off() {
	l.displayMode.On = false
	l.execInstruction(insSetDisplayMode, l.displayMode.byte())
}

// SetBlink will enable/disable blinking of the cursor
func (l *LCD) SetBlink(enabled bool) {
	l.displayMode.Blink = enabled
	l.execInstruction(insSetDisplayMode, l.displayMode.byte())
}

// SetCursor will enable/disable the cursor
func (l *LCD) SetCursor(enabled bool) {
	l.displayMode.Cursor = true
	l.execInstruction(insSetDisplayMode, l.displayMode.byte())
}
