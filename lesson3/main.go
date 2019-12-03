package main

import (
	"machine"
	"time"
)

//
// Take input from an analog sensor
//

func main() {
	machine.InitADC() // init the machine's ADC subsystem
	machine.InitPWM() // init the machine's PWM subsystem

	rotarySensor := machine.ADC{machine.A0}
	machine.UART1.Configure(machine.UARTConfig{TX: machine.PA22, RX: machine.PA23})
	rotarySensor.Configure()

	led := machine.PWM{machine.D11}
	led.Configure()

	for {
		sensorValue := rotarySensor.Get()
		println("Value: ", sensorValue)
		led.Set(sensorValue)
		time.Sleep(time.Millisecond * 10)
	}
}
