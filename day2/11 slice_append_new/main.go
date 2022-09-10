package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	a2 := a1[:]

	a2 = append(a1[:1], a1[2:]...)
	fmt.Println(a2)
	fmt.Println(a1)
}
