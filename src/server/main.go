package main

import (
	"flag"
	"fmt"
)


func main() {
	port := flag.Int("port", DEFAULT_PORT, "")
	filedir := flag.String("filedir", default_dir(), "")
	flag.Parse()

	fmt.Println("Starting server at port", *port)
	start(*port, *filedir)
}
