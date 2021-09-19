package lcd

type entryMode struct {
	Increment bool
	Shift     bool
}

func (e entryMode) byte() byte {
	var output byte
	if e.Increment {
		output |= 0b10
	}
	if e.Shift {
		output |= 0b1
	}
	return output
}

// SetCursorIncrement controls cursor direction - increment/decrement i.e. right/left
func (l *LCD) SetCursorIncrement(enabled bool) {
	l.entryMode.Increment = enabled
	l.execInstruction(insEntryModeSet, l.entryMode.byte())
}

// SetShift controls content shifting - content will be moved to the left/right depending
// on the cursor incrementation setting
func (l *LCD) SetShift(enabled bool) {
	l.entryMode.Shift = enabled
	l.execInstruction(insEntryModeSet, l.entryMode.byte())
}
