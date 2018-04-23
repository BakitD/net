package main


import (
	"net"
	"fmt"
)


const (
	MESSAGE_TO_SERVER = "This message should be sent to server!"
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

func Start(data []byte) {

	for _, server := range SERVERS {
		conn, err := connect(server)
		if err == nil {
			conn.Write(data)
			conn.Close()
		} else {
			fmt.Println(err)
		}

	}
}

func main() {
	Start([]byte(MESSAGE_TO_SERVER))
}