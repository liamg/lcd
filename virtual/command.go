package virtual

import "time"

const (
	insClearDisplay byte = 1 << iota
	insReturnHome
	insEntryModeSet
	insSetDisplayMode
	insSetCursorDisplayShift
	insSetFunction
	insSetCGRAMAddress
	insSetDDRAMAddress
)

func (l *LCD) processInstruction() {
	instruction := l.read()

	switch {
	case instruction&insSetFunction > 0:
		l.setFunc(instruction)
		return
	}

	l.initStage = 0
}

func (l *LCD) setFunc(data byte) {
	if data&0b10000 > 0 {
		switch l.initStage {
		case 0:
			l.initStage++
			l.lastInitStage = time.Now()
			return
		case 1:
			if time.Since(l.lastInitStage) >= 41*(time.Millisecond/10) {
				l.initStage++
				l.lastInitStage = time.Now()
			}
			return
		case 2:
			if time.Since(l.lastInitStage) >= 100*time.Microsecond {
				l.initStage++
				l.lastInitStage = time.Now()
			}
			return
		case 3:
			l.initStage = 0
			l.initialised = true
		}
	}

	// TODO process "real" function set
}
