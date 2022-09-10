package main

import "fmt"

func main() {
	//fmt占位符 %s %d %x %o %b %T %v

	i := 2
	fmt.Printf("%T\t", i)
	fmt.Printf("%v\t", i)
	fmt.Printf("%b\t", i)
	fmt.Printf("%d\t", i)
	fmt.Printf("%o\t", i)
	fmt.Printf("%x\t", i)

	s := "linjianshu"
	fmt.Printf("%s\t", s)
	fmt.Printf("%v\t", s)
	fmt.Printf("%#v\t", s)
}
