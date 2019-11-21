package main

import (
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/espat"
	"tinygo.org/x/drivers/espat/mqtt"
	"tinygo.org/x/drivers/ssd1306"
)

var (
	buttonPush bool

	uart = machine.UART1
	tx   = machine.PA22
	rx   = machine.PA23

	adaptor *espat.Device
	topic   = "tinygo"

	display ssd1306.Device
)

// access point info. Change this to match your WiFi connection information.
const ssid = "arduino"
const pass = "ts2019"

// IP address of the MQTT broker to use. Replace with your own info, if so desired.
const server = "tcp://192.168.1.8:1883"

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	machine.InitADC()
	machine.InitPWM()

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	blue := machine.LED
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Init esp8266/esp32
	adaptor = espat.New(uart)
	adaptor.Configure()

	// first check if connected
	if connectToESP() {
		blue.High()
		println("Connected to wifi adaptor.")
		adaptor.Echo(false)

		blue.Low()
		connectToAP()
		blue.High()
	} else {
		println("")
		failMessage("Unable to connect to wifi adaptor.")
		return
	}

	opts := mqtt.NewClientOptions(adaptor)
	opts.AddBroker(server).SetClientID("tinygo-client-" + "unit1")

	blue.Low()
	println("Connectng to MQTT...")
	cl := mqtt.NewClient(opts)
	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		failMessage(token.Error().Error())
	}

	for {
		buttonPush = button.Get()
		if !buttonPush {
			blue.Low()
		} else {
			blue.High()
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

	// Right now this code is only reached when there is an error. Need a way to trigger clean exit.
	println("Disconnecting MQTT...")
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
	println(adaptor.GetClientIP())
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

