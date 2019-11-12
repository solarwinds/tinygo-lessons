package main

import(
	"machine"
	"time"

	"tinygo.org/x/drivers/espat"
	"tinygo.org/x/drivers/espat/mqtt"
)

//
// Send data to MQTT
//

var (
	uart = machine.UART1
	tx = machine.PA22  // ARV TX pin
	rx = machine.PA23  // ARV RX pin

	wifiAdapter *espat.Device
)

const(
	wifiSSID   = "arduino"
	wifiPass   = "tinygo123"
)

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D12
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

  if connected := connectedToWiFi(); !connected{
    println("Can't connect to Wifi")
    return
  }

  connectToAP()

	wifiAdapter = espat.New(uart)
	wifiAdapter.Configure()

  client := mqttClient(wifiAdapter, "192.168.1.8:1883", "tiny-go")

  if token := client.Connect(); token.Wait() && token.Error() != nil {
    println("Error: ", token.Error().Error())
  }

	for {
		if button.Get(){
      println("Publishing MQTT message")
      token := publishMqttMessage(client, "tinygo-lessons", "hello, world")

      token.Wait()

      if token.Error() != nil{
        println(token.Error().Error())
      }

			led.High()
		}else{
			led.Low()
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func publishMqttMessage(c mqtt.Client, topic, msg string) mqtt.Token {
  data := "{\"msg:\"\"" + msg + "\"}"
  return c.Publish(topic, 0, false, []byte(data))
}

// mqttClient returns a client for talking to an MQTT broker
func mqttClient(e *espat.Device, server, clientID string) mqtt.Client{
  opts := mqtt.NewClientOptions(e)
  opts.AddBroker(server).SetClientID(clientID)
  return mqtt.NewClient(opts)
}


func connectedToWiFi() bool {
	for	i := 0; i < 10; i++ {
		println("Connecting to ESP wifi adapter")
		if wifiAdapter.Connected() {
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// connect to access point
func connectToAP() {
	println("Connecting to wifi network...")

	wifiAdapter.SetWifiMode(espat.WifiModeClient)
	wifiAdapter.ConnectToAP(wifiSSID, wifiPass, 10)

	println("Connected.")
	println(wifiAdapter.GetClientIP())
}


