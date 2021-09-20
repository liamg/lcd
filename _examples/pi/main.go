package main

import (
	"time"

	"github.com/liamg/lcd"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

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

	lcd.WriteLine(0, "Time:")
	for {
		lcd.WriteLine(1, time.Now().Format("15:04:05"))
		time.Sleep(time.Second)
	}
}
