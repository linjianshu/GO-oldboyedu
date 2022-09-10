package main

import (
	"fmt"
	"strings"
)

var m = make([]string, 4)

//Split 切割字符串 a b c => [a c]
func Split(str, sep string) []string {
	i := 0
	judge(str, sep, i)
	return m
}

func judge(s string, sep string, i int) {
	if !strings.Contains(s, sep) {
		return
	}

	index := strings.Index(s, sep)
	if s[:index] != "" {
		m[i] = s[:index]
	} else {
		m[i] = s[index+1:]
	}
	if s[index+1:] != "" {
		m[i+1] = s[index+1:]
	}
	judge(m[i+1], sep, i+1)
}

func main() {
	got := Split("abcb", "b")
	fmt.Println(got)
}
