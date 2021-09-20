package lcd

import "github.com/stianeikeland/go-rpio/v4"

type piPin struct {
	inner rpio.Pin
}

func PiPin(p uint8) Pin {
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
