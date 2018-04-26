package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"../utils"
)

func format_address(port int) string {
	return ":" + strconv.Itoa(port)
}

func default_dir() string {
	cwd, _ := os.Getwd()
	abspath, _ := filepath.Abs(filepath.Join(cwd, DEFAULT_DIRECTORY))
	return abspath
}

func pick_file(filename string) string {
	path, _ := filepath.Abs(filepath.Join(default_dir(), filename))
	return path
}

func send_file(filename string, conn net.Conn) int {
	var bytes_sent int = 0
	var buffer []byte = make([]byte, BUFFER_SIZE)

	f, err := os.Open(pick_file(filename))
	if err != nil {
		utils.Print_error(err)
		return bytes_sent
	}

	reader := bufio.NewReaderSize(f, BUFFER_SIZE)
	for {
		bytes_read, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			utils.Print_error(err)
			break
		}
		if bytes_read > 0 {
			n, err := conn.Write(buffer[:bytes_read])
			if err != nil {
				utils.Print_error(err)
				break
			}
			bytes_sent += n
		} else {
			break
		}
	}
	return bytes_sent
}

func handle_connection(conn net.Conn) int {
	var buffer []byte = make([]byte, FILENAME_BUFFER_MAX_LENGTH)
	var bytes_sent int = 0

	defer conn.Close()

	bytes_read, err := conn.Read(buffer)
	if err != nil {
		utils.Print_error(err)
		return bytes_sent
	}
	if bytes_read == 0 {
		utils.Print_message("Received filename with zero length")
		return bytes_sent
	}
	return send_file(string(buffer[:bytes_read]), conn)
}

func start(port int, filedir string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", format_address(port))
	utils.Handle_error(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.Handle_error(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		} else {
			sent := handle_connection(conn)
			utils.Print_message(fmt.Sprintf("%d bytes were sent.", sent))
		}
	}
}
