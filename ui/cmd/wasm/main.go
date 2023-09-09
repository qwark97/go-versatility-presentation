package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("readTemperature", readTemperature())
	select {}
}

func calculateTemp() (string, error) {
	temps := []string{"21.3°C", "22°C", "21.6°C"}
	idx := rand.Int() % 3
	return string(temps[idx]), nil
}

func readTemperature() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 0 {
			return "Invalid no of arguments passed"
		}
		temp, err := calculateTemp()
		if err != nil {
			return err.Error()
		}

		document := js.Global().Get("document")
		if !document.Truthy() {
			return "Unable to get document object"
		}

		tempPlaceholder := document.Call("getElementById", "temperature")
		if !tempPlaceholder.Truthy() {
			return "Unable to get output text area"
		}

		fmt.Println("temp", temp)
		tempPlaceholder.Set("innerHTML", temp)

		return temp
	})

	return jsonfunc
}
