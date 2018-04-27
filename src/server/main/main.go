package main

import (
	"flag"
	"fmt"
	"server"
)

func main() {
	port := flag.Int("port", server.DEFAULT_PORT, "")
	filedir := flag.String("filedir", server.Default_dir(), "")
	flag.Parse()

	fmt.Println("Starting server at port", *port)
	server.Start(*port, *filedir)
}
