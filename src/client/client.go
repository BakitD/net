package main

import (
	"fmt"
	"net"
)

var SERVERS []string = []string{":9999"}

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

func start(data string) {
	var buffer []byte = make([]byte, 512)

	for _, server := range SERVERS {
		conn, err := connect(server)
		if err == nil {
			conn.Write([]byte(data))
			for {
				bytes_read, _ := conn.Read(buffer)
				if bytes_read == 0 {
					break
				} else {
					fmt.Println(string(buffer[:bytes_read]))
				}
			}
		} else {
			fmt.Println(err)
		}

	}
}