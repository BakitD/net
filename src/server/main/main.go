package main

import (
	"flag"
	"log"
	"server"
)

func main() {
	port := flag.Int("port", server.DEFAULT_PORT, "")
	filedir := flag.String("filedir", server.Default_dir(), "")
	flag.Parse()

	log.Print("Starting server at port ", *port)
	server.Start(*port, *filedir)
}
