# lcd1602

A Go module for driving LCD1602 devices.

![Demo](demo.png)

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


