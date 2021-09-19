package main

import (
	"time"

	"github.com/liamg/lcd1602"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

	lcd, err := lcd1602.Connect(
		rpio.Pin(24),                                         // RS
		rpio.Pin(25),                                         // E
		rpio.Pin(5), rpio.Pin(6), rpio.Pin(13), rpio.Pin(19), // DATA
	)
	if err != nil {
		panic(err)
	}

	lcd.WriteTopLine("Time:")
	for {
		lcd.WriteBottomLine(time.Now().Format("15:04:05"))
		time.Sleep(time.Second)
	}
}
