package main

import (
	"client"
	"os"
)

func main() {
	source := os.Args[1]
	args := os.Args[2:]

	for _, filename := range args {
		client.Start(source, filename)
	}
}
