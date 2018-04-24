package main


import (
	"net"
	"fmt"
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

	for _, server := range SERVERS {
		conn, err := connect(server)
		if err == nil {
			conn.Write([]byte(data))
			conn.Close()
		} else {
			fmt.Println(err)
		}

	}
}
