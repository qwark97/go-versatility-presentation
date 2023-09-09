package presentation

import (
	"encoding/json"
	"fmt"
)

type Stdout struct {
}

func NewStdout() *Stdout {
	return &Stdout{}
}

func (p *Stdout) Show(data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}
