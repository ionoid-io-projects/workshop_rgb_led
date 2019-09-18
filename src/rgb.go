/*

A blinker example using go-rpio library.
Requires administrator rights to run

Toggles a LED on physical pin 19 (mcu pin 10)
Connect a LED with resistor from pin 19 to ground.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pinRed   = rpio.Pin(25)
	pinGreen = rpio.Pin(24)
	pinBlue  = rpio.Pin(23)
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pinRed.Output()
	pinGreen.Output()
	pinBlue.Output()

	Red()
	time.Sleep(time.Second * 2)
	Init()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		color, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		color = strings.TrimSuffix(color, "\n")

		switch color {
		case "red":
			Red()
		case "green":
			Green()
		case "orange":
			Orange()
		default:
			Init()

		}
	}
}

func Red() {
	pinRed.High()
	pinGreen.Low()
	pinBlue.Low()
}

func Green() {
	pinRed.Low()
	pinGreen.High()
	pinBlue.Low()
}

func Orange() {
	pinRed.High()
	pinGreen.Low()
	pinBlue.High()
}

func Init() {
	pinRed.Low()
	pinGreen.Low()
	pinBlue.Low()
}
