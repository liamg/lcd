package main

import (
	"time"

	"github.com/liamg/lcd1602"
)

func main() {
	lcd, err := lcd1602.Connect(24, 12, 25, 5, 6, 13, 19)
	if err != nil {
		panic(err)
	}
	for {
		lcd.Clear()
		lcd.WriteTopLine(time.Now().Format("15:04:05"))
		time.Sleep(time.Second)
	}
	lcd.Close()
}
