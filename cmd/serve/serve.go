package main

import (
	"log"

	"github.com/jesses-code-adventures/tiver/server"
)

func run() (err error) {
	server := server.NewServer()
	err = server.ListenAndServe()
	return
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
