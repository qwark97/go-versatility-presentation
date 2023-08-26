package flags

import (
	"flag"
	"time"
)

const (
	defaultAddr       = "localhost:8080"
	defaultReqTimeout = time.Duration(10 * time.Second)
	defaultConfPath   = "./.hub/peripherals.json"
)

type Configuration struct {
	addr       string
	reqTimeout time.Duration
	confPath   string
}

func Parse() Configuration {
	c := Configuration{}
	flag.StringVar(&c.addr, "addr", defaultAddr, "defines address of the HTTP server including PORT")
	flag.DurationVar(&c.reqTimeout, "reqTimeout", defaultReqTimeout, "defines maximum time of each request")
	flag.StringVar(&c.confPath, "confPath", defaultConfPath, "defines path to the peripherals configurations")

	flag.Parse()
	return c
}

func (c Configuration) Addr() string {
	return c.addr
}

func (c Configuration) RequestTimeout() time.Duration {
	return c.reqTimeout
}

func (c Configuration) Path() string {
	return c.confPath
}
