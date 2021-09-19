package virtual

import "time"

type LCD struct {
	initialised    bool
	initStage      uint8
	lastInitStage  time.Time
	ePin           *Pin
	enableHighTime time.Time
	rsPin          *Pin
	dataPins       []*Pin
}

func New(columns, rows uint8) *LCD {
	device := &LCD{}
	device.ePin = newPin(device.enableChange)
	device.rsPin = newPin(nil)
	for i := 0; i < 8; i++ {
		device.dataPins = append(device.dataPins, newPin(nil))
	}
	return device
}

func (l *LCD) IsInitialised() bool {
	return l.initialised
}

func (l *LCD) EnablePin() *Pin {
	return l.ePin
}

func (l *LCD) RegisterSelectPin() *Pin {
	return l.rsPin
}

func (l *LCD) DataPins() []*Pin {
	return l.dataPins
}

func (l *LCD) enableChange(output, state bool) {
	if !output {
		return
	}
	if state {
		l.enableHighTime = time.Now()
		return
	}

	if time.Since(l.enableHighTime) < time.Microsecond {
		return
	}

	if l.rsPin.state {
		l.processData()
		return
	}

	l.processInstruction()
}
