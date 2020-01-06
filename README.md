# golang-raspberry-pi-gpio-test

Simple demo test project to learn golang and gpio.

## Build for raspberry pi

````shell script
env GOOS=linux GOARCH=arm GOARM=5 go build -o gpio-test
````

## Install on raspberry pi

````shell script
scp gpio-test pi@raspberrypi.local:/home/pi/test
````