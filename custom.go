package lcd

import (
	"fmt"
)

func (l *LCD) DefineCharacter(id uint8, data [8]byte) error {
	if id > 7 {
		return fmt.Errorf("character id must be in the range 0-7")
	}

	l.setCGRAMAddress(id * 0x8)

	for _, row := range data {
		l.Write(row & 0b00011111)
	}

	return nil
}
