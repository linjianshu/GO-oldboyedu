package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%#v\n", os.Args[2])
	fmt.Printf("%T\n", os.Args)
}
