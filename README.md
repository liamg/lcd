# lcd1602

A Go module for driving common LCD devices (those using the HD44780 controller or similar.)

Built for Raspberry Pi, but it should work with any other device where you can use an implementation of the `lcd.Pin` interface.

## Example

```golang
package main

import (
	"github.com/liamg/lcd"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	_ = rpio.Open()
	defer rpio.Close()

	lcd, _ := lcd.New1602(
		rpio.Pin(24),                                         // RS
		rpio.Pin(25),                                         // E
		rpio.Pin(5), rpio.Pin(6), rpio.Pin(13), rpio.Pin(19), // DATA
	)

	lcd.WriteLine(0, "Hello World!")
	lcd.WriteLine(1, ":)")
}
```

![Demo](demo.png)

## TODO

- Add support for custom characters

