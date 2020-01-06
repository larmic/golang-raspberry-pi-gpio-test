package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"os/signal"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		exit()
		os.Exit(1)
	}()

	pin.Input()

	for {
		fmt.Println("Pin value: ", pin.Read())
		time.Sleep(time.Millisecond * 10)
	}
}

func exit() int {
	defer fmt.Println("Unmap gpio memory!")
	defer rpio.Close()

	return 3
}
