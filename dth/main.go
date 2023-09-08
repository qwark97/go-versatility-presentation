package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
	"tinygo.org/x/drivers/net/http"
	"tinygo.org/x/drivers/wifinina"
)

func main() {
	pin := machine.D5

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

	device := dht.New(pin, dht.DHT11)

	http.UseDriver(adaptor)
	http.HandleFunc("/temp", func(w http.ResponseWriter, _ *http.Request) {
		device.ReadMeasurements()
		temp, _ := device.Temperature()
		fmt.Fprintf(w, fmt.Sprintf("%02d.%d", temp/10, temp%10))
	})
	http.HandleFunc("/humi", func(w http.ResponseWriter, _ *http.Request) {
		device.ReadMeasurements()
		humi, _ := device.Humidity()
		fmt.Fprintf(w, fmt.Sprintf("%02d.%d", humi/10, humi%10))
	})

	http.ListenAndServe(":80", nil)
}

func connectToAP(adaptor *wifinina.Device) {
	var (
		ssid = "ssid"
		pass = "pass"
	)
	time.Sleep(2 * time.Second)
	for {
		println("Connecting to " + ssid)
		err := adaptor.ConnectToAccessPoint(ssid, pass, 10*time.Second)
		if err != nil {
			println("err:", err)
			continue
		}
		println("Connected.")
		return
	}
}
