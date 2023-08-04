package main

import (
	"log"

	cmd "github.com/omani/nkn-openapi-client/cmd/commands"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Panic: %+v", r)
		}
	}()

	cmd.Execute()
}
