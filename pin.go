package lcd

type Pin interface {
	High()
	Low()
	Output()
	Input()
	Read() bool
}

type pins struct {
	registerSelect Pin
	readWrite      Pin
	enable         Pin
	data           []Pin
}

func (p *pins) init() {
	p.registerSelect.Output()
	p.enable.Output()
	if p.readWrite != nil {
		p.readWrite.Output()
	}
	p.setDataOutput()
}

func (p *pins) setDataInput() {
	for _, data := range p.data {
		data.Input()
	}

}

func (p *pins) setDataOutput() {
	for _, data := range p.data {
		data.Output()
	}
}
