package main

import (
	"log"
	"os"
)

func main() {
	cli := newCLI()

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
