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

	rw := rpio.Pin(12)
	rw.Output()
	rw.Low()

	lcd, err := lcd.New(
		16,
		2,
		rpio.Pin(24),                                         // RS
		rpio.Pin(25),                                         // E
		rpio.Pin(5), rpio.Pin(6), rpio.Pin(13), rpio.Pin(19), // DATA
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
