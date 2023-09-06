package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/net/http"
	"tinygo.org/x/drivers/wifinina"
)

func main() {
	var led = machine.LED
	led.Configure(
		machine.PinConfig{
			Mode: machine.PinOutput,
		},
	)

	spi := machine.NINA_SPI
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	adaptor := wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()

	connectToAP(adaptor)
	displayIP(adaptor)

	http.UseDriver(adaptor)

	http.HandleFunc("/off", func(w http.ResponseWriter, _ *http.Request) {
		led.Low()
		println("off")
		fmt.Fprintf(w, "led.High()")
	})
	http.HandleFunc("/on", func(w http.ResponseWriter, _ *http.Request) {
		led.High()
		println("on")
		fmt.Fprintf(w, "led.High()")
	})

	http.ListenAndServe(":80", nil)
}

// connect to access point
func connectToAP(adaptor *wifinina.Device) {
	var (
		ssid = "ssid"
		pass = "pass"
	)
	time.Sleep(2 * time.Second)
	var err error
	for i := 0; i < 3; i++ {
		println("Connecting to " + ssid)
		err = adaptor.ConnectToAccessPoint(ssid, pass, 10*time.Second)
		if err == nil {
			println("Connected.")
			return
		}
	}
	println(err.Error())
}

func displayIP(adaptor *wifinina.Device) {
	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		println(err.Error())
		time.Sleep(1 * time.Second)
	}
	println("IP address: " + ip.String())
}
