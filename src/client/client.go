package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"../utils"
)

var SERVERS []string = []string{":9999"}


func default_dir() string {
	cwd, _ := os.Getwd()
	abspath, _ := filepath.Abs(filepath.Join(cwd, DEFAULT_SAVE_DIRECTORY))
	return abspath
}

func connect(address string) (conn *net.TCPConn, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func find_file(filename string, conn net.Conn) (int, error) {
	return conn.Write([]byte(filename))
}

func receive_file(filename string, conn net.Conn) (int, error) {
	var bytes_written int = 0
	var buffer []byte = make([]byte, BUFFER_LENGTH)

	path := filepath.Join(default_dir(), filepath.Base(filename))
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return bytes_written, err
	}

	defer f.Close()

	for {
		bytes_read, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			return bytes_written, err
		}
		if bytes_read == 0 {
			break
		} else {
			n, err := f.Write(buffer[:bytes_read])
			if err != nil {
				return bytes_written + n, err
			}
			bytes_written += n
		}
	}
	return bytes_written, nil
}

func start(filename string) {
	for _, server := range SERVERS {
		conn, err := connect(server)
		if err == nil {
			_, err = find_file(filename, conn)
			if err != nil {
				utils.Print_error(err)
			} else {
				bytes_received, err := receive_file(filename, conn)
				if err != nil {
					utils.Print_error(err)
				} else {
					utils.Print_message(fmt.Sprintf("%d bytes were received", bytes_received))
					break
				}
			}
		} else {
			utils.Print_error(err)
		}
	}
}
