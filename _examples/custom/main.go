package main

import (
	"github.com/liamg/lcd"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

	rw := lcd.PiPin(12)
	rw.Output()
	rw.Low()

	lcd, err := lcd.New1602(
		lcd.FontSize5x8,
		lcd.PiPin(24), // RS
		lcd.PiPin(25), // E
		nil,
		lcd.PiPin(5), lcd.PiPin(6), lcd.PiPin(13), lcd.PiPin(19), // DATA
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
