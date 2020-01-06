package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"time"
)

var (
	// Use mcu pin 27, corresponds to physical pin 13 on the pi
	// see https://pinout.xyz/pinout/pin13_gpio27
	pin = rpio.Pin(27)
)

func main() {
	fmt.Println("Read all 100ms PIN 27 (physical 13)")

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	pin.Input()

	// Toggle pin 20 times
	for x := 0; x < 200; x++ {
		fmt.Println(pin.Read())
		time.Sleep(time.Second / 10)
	}
}
