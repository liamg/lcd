package lcd

import (
	"sync"

	"github.com/stianeikeland/go-rpio/v4"
)

type piPin struct {
	inner rpio.Pin
}

var gpioLock sync.Mutex
var gpioOpen bool

func openGPIO() {
	gpioLock.Lock()
	defer gpioLock.Unlock()
	if gpioOpen {
		return
	}
	_ = rpio.Open()
	gpioOpen = true
}

func closeGPIO() {
	gpioLock.Lock()
	defer gpioLock.Unlock()
	if !gpioOpen {
		return
	}
	gpioOpen = false
	_ = rpio.Close()
}

func PiPin(p uint8) Pin {
	openGPIO()
	return &piPin{
		inner: rpio.Pin(p),
	}
}

func (p *piPin) Read() bool {
	return p.inner.Read() == rpio.High
}

func (p *piPin) High() {
	p.inner.High()
}

func (p *piPin) Low() {
	p.inner.Low()
}

func (p *piPin) Input() {
	p.inner.Input()
}

func (p *piPin) Output() {
	p.inner.Output()
}
