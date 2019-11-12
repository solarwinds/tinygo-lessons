package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

var(
	msgChan = make(chan [2]string)
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("localhost:1883")
	opts.SetClientID("tinygo-lessons")
	opts.SetDefaultPublishHandler(handler)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("tinygo", 0, nil); token.Wait() && token.Error() != nil{
		log.Fatalln(token.Error())
	}

	fmt.Println("starting infinite consumer")

	for {
		incoming := <- msgChan
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
	}
}

func handler(c mqtt.Client, msg mqtt.Message)  {
	msgChan <- [2]string{msg.Topic(), string(msg.Payload())}
}
