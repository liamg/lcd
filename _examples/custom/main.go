package main

import (
	"github.com/liamg/lcd"
)

func main() {

	lcd, err := lcd.New1602(
		lcd.FontSize5x8,
		lcd.PiPin(24), // RS
		lcd.PiPin(25), // E
		lcd.PiPin(12), // RW (set this to nil if you don't want to use it)
		lcd.PiPin(5),  // DB4
		lcd.PiPin(6),  // DB5
		lcd.PiPin(13), // DB6
		lcd.PiPin(19), // DB7
	)
	if err != nil {
		panic(err)
	}

	lcd.Clear()

	lcd.DefineCharacter(0, [8]byte{
		0b00000,
		0b11011,
		0b11011,
		0b00000,
		0b10001,
		0b10001,
		0b01110,
		0b00000,
	})

	lcd.WriteLine(0, "Custom: \x00")
}
