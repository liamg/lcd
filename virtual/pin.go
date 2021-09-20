package virtual

type Pin struct {
	state    bool
	output   bool
	onChange func(output, state bool)
}

func newPin(onChange func(output, state bool)) *Pin {
	return &Pin{
		onChange: onChange,
	}
}

func (p *Pin) Output() {
	p.output = true
	p.change()
}

func (p *Pin) High() {
	p.state = true
	p.change()
}

func (p *Pin) Low() {
	p.state = false
	p.change()
}

func (p *Pin) Input() {
	p.output = false
	p.change()
}

func (p *Pin) Read() bool {
	return p.state
}

func (p *Pin) change() {
	if p.onChange == nil {
		return
	}
	p.onChange(p.output, p.state)
}
