package main

import "github.com/liamg/lcd1602"

func main() {
	lcd, err := lcd1602.Connect(24, 12, 25, 5, 6, 13, 19)
	if err != nil {
		panic(err)
	}
	defer lcd.Close()
	lcd.WriteLines("blah")
}
