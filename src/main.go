package main


import (
	"os"
	"fmt"
	"server"
)

const (
	DEFAULT_PORT = "9999"
)


func main() {
	var port string

	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = DEFAULT_PORT
	}

	fmt.Println("Starting at port ", port)
	server.Start(port)
}