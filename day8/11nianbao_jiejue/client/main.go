package main

import (
	"fmt"
	"net"
	"src/code.oldboyedu.com/day8/11nianbao_jiejue/protocol"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed , error : ", err)
		return
	}
	defer dial.Close()

	for i := 0; i < 20; i++ {
		encode, err := protocol.Encode("how are you , hello!")
		if err != nil {
			return
		}
		dial.Write(encode)
		//dial.Write([]byte("how are you , hello !"))
	}
}
