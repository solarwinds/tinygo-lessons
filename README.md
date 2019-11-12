# tinygo-lessons
Using TinyGo to program the Arduino Nano 33 IoT

* Setup [#setup]
* Lesson Structure [#lesson-structure]

## Setup
### Go
TinyGo uses Go to build. And Go's a great thing to have around anyway.
#### Linux
```
sudo apt update
sudo apt-get install golang-1.13-go
```

#### macOS
```
brew update
brew install go
```

#### Windows (Experimental)
Install from MSI or source https://golang.org/dl/

### TinyGo
Use the links below to understand how to get TinyGo installed on your system.
#### Linux
https://tinygo.org/getting-started/linux/

#### macOS
https://tinygo.org/getting-started/macos/

#### Windows (Experimental)
https://tinygo.org/getting-started/windows/

### TinyGo Drivers
```
go get -u tinygo.org/x/drivers
go get -u github.com/eclipse/paho.mqtt.golang
go get -u github.com/conejoninja/tinydraw
go get -u github.com/conejoninja/tinyfont
```

### Clang 8
http://releases.llvm.org/download.html#9.0.0

### bossac CLI tool
#### Linux
```
sudo apt install \
  libwxgtk3.0-dev
  libreadline-dev
git clone https://github.com/shumatech/BOSSA.git
cd BOSSA
make
```

### macOS or Windows
[Download the installer](https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-1.9.1.dmg)



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

### Docker-based MQTT Server
`docker run -it -p 1883:1883 -p 9001:9001 eclipse-mosquitto`

## Credits
Thanks to TinyGo creator Ayke van Laethem (@aykevl) and Ron Evans (@deadprogram) of the [HybridGroup](https://hybridgroup.com) for their work on TinyGo, related [workshop content](https://github.com/hybridgroup/hacklab-2019), and valuable input into these lessons.
