package virtual

func (l *LCD) processData() {

}

func (l *LCD) read() byte {
	var output byte
	for i := 0; i < 8; i++ {
		if l.dataPins[i].state {
			output |= 1 << i
		}
	}
	return output
}
