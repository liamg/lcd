package lcd

type FontSize uint8

const (
	FontSize5x8 FontSize = iota
	FontSize5x11
)

func (f FontSize) String() string {
	switch f {
	case FontSize5x8:
		return "5x8"
	case FontSize5x11:
		return "5x11"
	default:
		return "?"
	}
}

func (f FontSize) Width() uint8 {
	return 5
}

func (f FontSize) Height() uint8 {
	switch f {
	case FontSize5x8:
		return 8
	case FontSize5x11:
		return 11
	default:
		return 0
	}
}
