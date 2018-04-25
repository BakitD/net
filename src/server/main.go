package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)


func default_dir() string {
	cwd, _ := os.Getwd()
	abspath, _ := filepath.Abs(filepath.Join(cwd, DEFAULT_DIRECTORY))
	return abspath
}


func main() {
	port := flag.Int("port", DEFAULT_PORT, "")
	filedir := flag.String("filedir", default_dir(), "")
	flag.Parse()

	fmt.Println("Starting server at port", *port)
	start(*port, *filedir)
}
