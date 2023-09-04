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
	res, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
