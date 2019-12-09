package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/espat"
	"tinygo.org/x/drivers/espat/mqtt"
)

var (
	buttonPush bool

	uart = machine.UART1
	tx   = machine.PA22
	rx   = machine.PA23

	adaptor *espat.Device
	topic   = "tinygo"
)

// access point info. Change this to match your WiFi connection information.
const ssid = "Arduino"
const pass = "tinygo123"

// IP address of the MQTT broker to use. Replace with your own info, if so desired.
const server = "tcp://172.20.36.137:1883"

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	machine.InitADC()
	machine.InitPWM()

	button := machine.D12
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	led := machine.D11
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	adaptor = espat.New(uart)
	adaptor.Configure()

	// first check if connected
	if connectToESP() {
		led.High()
		println("Connected to wifi adaptor.")
		adaptor.Echo(false)

		led.Low()
		connectToAP()
		led.High()
	} else {
		println("")
		println("Unable to connect to wifi adaptor.")
		return
	}

	opts := mqtt.NewClientOptions(adaptor)
	opts.AddBroker(server).SetClientID("tinygo-client-" + "unit1")

	led.Low()
	println("Connectng to MQTT...")
	cl := mqtt.NewClient(opts)
	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		println(token.Error().Error())
	}

	for {
		buttonPush = button.Get()
		if !buttonPush {
			led.Low()
		} else {
			led.High()
			println("Publishing MQTT message...")
			data := []byte("{\"e\":[{ \"n\":\"hello\", \"sv\":\"from unit1\" }]}")
			token := cl.Publish(topic, 0, false, data)
			token.Wait()
			if token.Error() != nil {
				println(token.Error().Error())
			}
		}

		time.Sleep(time.Millisecond * 100)
	}

	println("Error: disconnecting MQTT...")
	cl.Disconnect(100)

	println("Done.")
}

// connect to ESP8266/ESP32
func connectToESP() bool {
	for i := 0; i < 5; i++ {
		println("Connecting to wifi adaptor...")
		if adaptor.Connected() {
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// connect to access point
func connectToAP() {
	println("Connecting to wifi network...")

	adaptor.SetWifiMode(espat.WifiModeClient)
	adaptor.ConnectToAP(ssid, pass, 10)

	println("Connected.")
	println("Waiting 10s before asking for client IP.")

	time.Sleep(10 * time.Second)
	println(adaptor.GetClientIP())
}
