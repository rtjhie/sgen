package main

import (
	"log"
	"os"

	"github.com/rtjhie/sgen"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	path := os.Args[1]
	return sgen.Generate(path, nil)
}
