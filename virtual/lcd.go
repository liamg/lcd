package virtual

import (
	"sync"
	"time"
)

type LCD struct {
	charWidth      uint8
	charHeight     uint8
	columns        uint8
	lines          uint8
	initialised    bool
	initStage      uint8
	lastSend       byte
	sendInProgress bool
	lastInitStage  time.Time
	ePin           *Pin
	enableHighTime time.Time
	rsPin          *Pin
	rwPin          *Pin
	dataPins       []*Pin
	config         config
	ddram          DDRAM
	cgrom          CGROM
	cgram          CGRAM
	currentAddress address
	busy           bool
	busyMu         sync.Mutex
}

type address struct {
	kind     addressKind
	location uint8
}

type addressKind uint8

const (
	addressKindDDRAM addressKind = iota
	addressKindCGRAM
)

type config struct {
	busSize    uint8
	lines      uint8
	fontHeight uint8 // 8 or 11
	increment  bool
	shift      bool
	on         bool
	cursor     bool
	blink      bool
}

var defaultConfig = config{
	busSize:    8, // always start in 8-bit mode
	lines:      1,
	fontHeight: 8,
	increment:  true,
}

func New(charWidth, charHeight, columns, lines uint8) *LCD {
	device := &LCD{
		charWidth:  charWidth,
		charHeight: charHeight,
		config:     defaultConfig,
		columns:    columns,
		lines:      lines,
	}
	if charWidth != 5 {
		panic("unsupported character width")
	}
	switch charHeight {
	case 8:
		device.cgrom = &CGROM5x8{}
	case 11:
		device.cgrom = &CGROM5x11{}
	default:
		panic("unsupported character height")
	}
	device.ePin = newPin(device.enableChange)
	device.rsPin = newPin(nil)
	device.rwPin = newPin(nil)
	for i := 0; i < 8; i++ {
		device.dataPins = append(device.dataPins, newPin(nil))
	}
	device.ddram.reset()
	return device
}

func (l *LCD) IsInitialised() bool {
	return l.initialised
}

func (l *LCD) BusSize() uint8 {
	return l.config.busSize
}

func (l *LCD) PhysicalSize() (cols uint8, lines uint8) {
	return l.columns, l.lines
}

func (l *LCD) MemoryLines() uint8 {
	return l.config.lines
}

func (l *LCD) FontSize() (width, height uint8) {
	return 5, l.config.fontHeight
}

func (l *LCD) EnablePin() *Pin {
	return l.ePin
}

func (l *LCD) ReadWritePin() *Pin {
	return l.rwPin
}

func (l *LCD) RegisterSelectPin() *Pin {
	return l.rsPin
}

func (l *LCD) DataPins() []*Pin {
	return l.dataPins
}

func (l *LCD) addressForLine(line uint8) byte {
	if line%2 == 0 {
		return 0x00 + (l.columns * (line / 2))
	}
	return 0x40 + (l.columns * ((line - 1) / 2))
}

func (l *LCD) String() string {
	var output string
	for line := byte(0); line < l.lines; line++ {
		if line > 0 {
			output += "\n"
		}
		address := l.addressForLine(line)
		for col := byte(0); col < l.columns; col++ {
			data := l.ddram.read(address + col)
			if data < 8 {
				output += "?"
			} else {
				output += string(rune(data))
			}
		}
	}
	return output
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

	if l.config.busSize == 4 {
		if !l.sendInProgress {
			l.lastSend = l.readPins()
			l.sendInProgress = true
			return
		}
		l.sendInProgress = false
		l.lastSend |= l.readPins() >> 4
	} else {
		l.lastSend = l.readPins()
	}

	if l.rsPin.state {
		l.processData(l.lastSend)
		return
	}

	l.processInstruction(l.lastSend)
}
