package flags

import "flag"

const (
	defaultAddr = "localhost:8080"
)

type Configuration struct {
	addr string
}

func Parse() Configuration {
	c := Configuration{}
	flag.StringVar(&c.addr, "addr", defaultAddr, "defines address of the HTTP server including PORT")

	flag.Parse()
	return c
}

func (c Configuration) Addr() string {
	return c.addr
}
