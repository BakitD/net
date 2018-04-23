package server

import (
	"os"
	"fmt"
	"net"
)


const (
	ERROR_MESSAGE = "Error %s"
	ERROR_CODE_AT_EXIT = 1
	BUFFER_SIZE = 256
)


func handle_error(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, ERROR_MESSAGE, e.Error())
		os.Exit(ERROR_CODE_AT_EXIT)
	}
}


func print_message(message []byte, n_bytes int) {
	fmt.Println("Received from client: ", string(message[:n_bytes]))
}


func format_address(port string) string {
	return ":" + port
}


func Start(port string) {
	var buffer []byte = make([]byte, BUFFER_SIZE)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", format_address(port))
	handle_error(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	handle_error(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		n_bytes, err := conn.Read(buffer)
		if err == nil {
			print_message(buffer, n_bytes)
		}
		conn.Close()
	}
}