package main

import (
	"fmt"
	"strings"
)

func main() {
	suffixFunc := makeSuffixFunc(".jpg")
	f := makeSuffixFunc(".txt")
	fmt.Println(suffixFunc("text"))
	fmt.Println(f("text"))
}

func makeSuffixFunc(suffix string) func(string2 string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
