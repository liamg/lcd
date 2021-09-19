package lcd

// see http://web.alfredstate.edu/faculty/weimandn/lcd/lcd_addressing/lcd_addressing_index.html
func (l *LCD) addressForLine(line uint8) byte {
	if line%2 == 0 {
		return 0x00 + (l.columns * (line / 2))
	}
	return 0x40 + (l.columns * ((line - 1) / 2))
}
