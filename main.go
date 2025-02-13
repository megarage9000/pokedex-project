package main

import (
	"time"
	"internal/pokeapi"
)

func main() {
	config := Configuration {
		client: pokeapi.NewClient(5 * time.Second, 1 * time.Minute),
	}
	startRepl(config)
}

