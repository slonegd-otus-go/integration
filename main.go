package main

import (
	"log"

	"github.com/slonegd-otus-go/integration/cmd"
)

func main() {
	if err := cmd.Command.Execute(); err != nil {
		log.Fatal(err)
	}
}
