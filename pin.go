package lcd1602

type Pin interface {
	High()
	Low()
	Output()
}

type pins struct {
	registerSelect Pin
	enable         Pin
	data           []Pin
}

func (p *pins) init() {
	p.registerSelect.Output()
	p.enable.Output()
	for _, data := range p.data {
		data.Output()
	}
}
