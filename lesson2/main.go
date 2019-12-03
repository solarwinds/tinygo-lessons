package main

import (
	"machine"
	"time"
)

//
// Take input from a machine pin
//
func main() {
	machine.UART1.Configure(machine.UARTConfig{TX: machine.PA22, RX: machine.PA23})

	led := machine.D11
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D12
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if button.Get() {
			led.High()
			println("^") // pew pew pew!
		} else {
			led.Low()
			println("")
		}
		time.Sleep(time.Millisecond * 10)
	}
}
