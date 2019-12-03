package main

import (
	"machine"
	"time"
	"tinygo.org/x/drivers/buzzer"
)

//
// Produce a buzz past a threshold
//
func main() {
	machine.InitADC() // init the machine's ADC subsystem
	machine.InitPWM() // init the machine's PWM subsystem

	rotarySensor := machine.ADC{machine.A0}
	machine.UART1.Configure(machine.UARTConfig{TX: machine.PA22, RX: machine.PA23})
	rotarySensor.Configure()

	led := machine.PWM{machine.D11}
	led.Configure()

	buzrPin := machine.D10
	buzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buzz := buzzer.New(buzrPin)

	for {
		sensorValue := rotarySensor.Get()

		if sensorValue > 32000 {
			buzz.On()
		} else {
			buzz.Off()
		}

		println("Value: ", sensorValue)
		led.Set(sensorValue)
		time.Sleep(time.Millisecond * 10)
	}
}
