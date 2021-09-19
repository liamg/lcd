# lcd1602

A Go module for driving LCD1602 devices.

Built for Raspberry Pi, but it should work with any other device where you can use an implementation of the `lcd1602.Pin` interface.

## Example

```golang
package main

import (
	"time"

	"github.com/liamg/lcd1602"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	_ = rpio.Open()
	defer rpio.Close()

	lcd, _ := lcd1602.Connect(
		rpio.Pin(24),                                         // RS
		rpio.Pin(25),                                         // E
		rpio.Pin(5), rpio.Pin(6), rpio.Pin(13), rpio.Pin(19), // DATA
	)

	lcd.WriteTopLine("Hello World!")
	lcd.WriteBottomLine(":)")
}
```

![Demo](demo.png)

