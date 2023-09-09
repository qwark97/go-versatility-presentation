package main

import (
	"fmt"
	"io"
	"net/http"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("readTemperature", readTemperature())
	select {}
}

var addr = "http://localhost:9001/api/v1/last-reading/db6c8320-8e65-4db7-b4f3-129bf22f0ee0"

func calculateTemp() (string, error) {
	client := http.DefaultClient

	resp, err := client.Get(addr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := string(b)
	return data, nil
}

func readTemperature() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 0 {
			return "Invalid no of arguments passed"
		}

		go func() {
			temp, err := calculateTemp()
			if err != nil {
				fmt.Println("err", err.Error())
				fmt.Println(err.Error())
			}

			document := js.Global().Get("document")
			if !document.Truthy() {
				fmt.Println("Unable to get document object")
			}

			tempPlaceholder := document.Call("getElementById", "temperature")
			if !tempPlaceholder.Truthy() {
				fmt.Println("Unable to get output text area")
			}

			fmt.Println("temp", temp)
			tempPlaceholder.Set("innerHTML", temp)
		}()

		return nil
	})

	return jsonfunc
}
