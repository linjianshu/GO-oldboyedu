package main

import "fmt"

const (
	da    = 1
	eat   = 2
	sleep = 4
)

func f(i int) {
	fmt.Printf("%b\n", i)
}

func main() {
	//二进制用途

	fmt.Printf("%b\n", da)
	fmt.Printf("%b\n", eat)
	fmt.Printf("%b\n", sleep)

	fmt.Printf("%b\n", da|eat)
}
