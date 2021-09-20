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
		rpio.Pin(24),                                         // RS
		rpio.Pin(25),                                         // E
		nil,                                                  // RW
		rpio.Pin(5), rpio.Pin(6), rpio.Pin(13), rpio.Pin(19), // DATA
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
