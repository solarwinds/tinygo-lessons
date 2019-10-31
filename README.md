# tinygo-lessons
Using TinyGo to program the Arduino Nano 33 IoT

* Setup [#setup]
* Lesson Structure [#lesson-structure]

## Setup

* LLVM
* Bossac

## Lesson Structure

(TODO: add macOS device callout here)

#### lesson0
"hello, world!" - blinking the onboard LED on an interval
	* Flashing basics

#### lesson1
Goroutines and UART
	* goroutines let you do more than one thing at once
	* TinyGo uses UART for built-in debugging b/c of its ubquity
	* `screen` to connect to serial interface

#### lesson2
Connecting an external LED to a pin
	* Breadboards and their rails

#### lesson3
Rotation sensor and more LED
	* Light up red or green based on rotation region
	* Send rotation info to UART for debugging

#### lesson4
Buzzer to create rotation threshold
	* Transforming analog values

#### lesson5
Sending data to an MQTT broker
	* Connect to wifi
	* Send MQTT message at rotation threshold


### LLVM
TinyGo uses LLVM's backend to target various architectures. You need to have at least LLVM 8 to get the [AVR backend](https://github.com/avr-llvm) support, which is necessary because [Arduino]() is [based on that architecture](https://en.wikipedia.org/wiki/AVR_microcontrollers).

## Credits
Thanks to TinyGo creator Ayke van Laethem (@aykevl) and Ron Evans (@deadprogram) of the [HybridGroup](https://hybridgroup.com) for their work on TinyGo, related [workshop content](https://github.com/hybridgroup/hacklab-2019), and valuable input into these lessons.
