package presentation

import (
	"fmt"
)

type Stdout struct {
}

func NewStdout() *Stdout {
	return &Stdout{}
}

func (p *Stdout) Show(data any) {
	fmt.Println(data)
}
