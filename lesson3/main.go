package main

import(
	"machine"
	"time"
)

//
// Take input from an analog sensor
//

const(
	ADC_REF    = float32(3.3) // reference voltage is 3.3v
	FULL_ANGLE = 300					// full angle value is 300 degrees
)

var (
	rotarySensor = machine.ADC{machine.A7}
)

func main() {
	machine.UART1.Configure(machine.UARTConfig{TX: machine.PA22, RX: machine.PA23})
	rotarySensor.Configure()
	//blue := machine.PWM{machine.D12}
	//blue.Configure()

	blue := machine.D12
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	blue.High()

	for {
		//blue.Set(rotarySensor.Get())
		println("Value: ", rotarySensor.Get())
		time.Sleep(time.Millisecond * 250)
	}
}


func angleFromValue(v uint16) float32{
	floatv := float32(v)
	voltage := float32((floatv * ADC_REF)/1023)
	println("Value: ", v)
	println("Voltage: ", voltage)
	degrees := (voltage * FULL_ANGLE) / ADC_REF
	return degrees
}




