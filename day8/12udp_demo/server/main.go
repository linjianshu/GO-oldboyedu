package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen failed , error : ", err)
		return
	}
	//不需要建立链接，直接收发数据
	var b [1024]byte
	defer udp.Close()
	for {
		n, addr, err := udp.ReadFromUDP(b[:])
		if err != nil {
			fmt.Println("read from udp failed , err: ", err)
			return
		}

		fmt.Println(b[:n])
		reply := strings.ToUpper(string(b[:n]))
		//发送数据
		udp.WriteToUDP([]byte(reply), addr)
	}

}
