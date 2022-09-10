package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//UDP client
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("dial failed , error : ", err)
		return
	}

	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	var reply [1024]byte
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		socket.Write([]byte(readString))
		//收回复的数据
		n, adder, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			return
		}
		fmt.Println("received data: ,addr :	", string(reply[:n]), adder)
	}

}
