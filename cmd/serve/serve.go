package main

import (
	"log"

	"github.com/jesses-code-adventures/tiver/env"
	"github.com/jesses-code-adventures/tiver/server"
)

func run() (err error) {
	s := server.NewServer()
	err = s.ListenAndServe()
	return
}

func main() {
	env.Load()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
