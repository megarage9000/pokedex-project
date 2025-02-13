// Separating commands into separate files
package main

import (
	"fmt"
	"os"
)

func commandExit(config *Configuration, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	config.client.StopCacheReap()
	defer os.Exit(0)
	return nil
}