package main

import (
	"github.com/tonsV2/todo-go/pgk/di"
	"github.com/tonsV2/todo-go/pgk/server"
	"log"
)

func main() {
	environment := di.GetEnvironment()

	e := server.GetEngine(environment)

	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
