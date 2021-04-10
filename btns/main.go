package main

import (
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bendahl/uinput"
	"github.com/stianeikeland/go-rpio"
)

var (
	green = rpio.Pin(23) // GPIO 4
	red = rpio.Pin(18) // GPIO 1
	backlight = rpio.Pin(24) // GPIO 5
)

func checkBtn(pin rpio.Pin, keyboard uinput.Keyboard, key int) {
	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge)

	for {
		if pin.EdgeDetected() {
			keyboard.KeyPress(key)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("btns"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer keyboard.Close()

	// Turn on backlight
	backlight.Output()
	backlight.High()

	// Check green button
	go checkBtn(green, keyboard, uinput.KeyK)

	// Check red button
	go checkBtn(red, keyboard, uinput.KeyN)

	// Wait until exit
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    <-c

	// Turn off backlight
	backlight.Low()
}

