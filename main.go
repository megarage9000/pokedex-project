package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	
	splits := strings.Fields(text)

	for i := 0; i < len(splits); i++ {
		splits[i] = strings.ToLower(splits[i])
	}

	return splits
}