package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	for _, filename := range args {
		start(filename)
	}
}
