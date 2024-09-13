package main

import (
	"log"

	"github.com/jesses-code-adventures/tiver/env"
	"github.com/jesses-code-adventures/tiver/sender"
)

func run() (err error) {
	s := sender.NewHttpSender()
	s.SendRequests()
	return
}

func main() {
	env.Load()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
