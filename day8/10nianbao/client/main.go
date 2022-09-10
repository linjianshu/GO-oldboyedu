package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed , error : ", err)
		return
	}
	defer dial.Close()

	for i := 0; i < 20; i++ {
		dial.Write([]byte("how are you , hello !"))
	}
}
