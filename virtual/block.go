package virtual

type block struct {
	w, h uint8
	data []bool
}

func newBlock(w, h uint8) *block {
	var b block
	b.w = w
	b.h = h
	b.clear()
	return &b
}

func (b *block) clear() {
	b.data = make([]bool, b.w*b.h)
}

func (b *block) set(x, y uint8, on bool) {
	offset := (int(y) * int(b.w)) + int(x)
	if offset >= len(b.data) {
		return
	}
	b.data[offset] = on
}

func (b *block) get(x, y uint8) bool {
	offset := (int(y) * int(b.w)) + int(x)
	if offset >= len(b.data) {
		return false
	}
	return b.data[offset]
}
