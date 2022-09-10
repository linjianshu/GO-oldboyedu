package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "hello沙河"
	sum := 0
	for _, v := range s1 {
		//if v >= 128 {
		//	sum++
		//}

		if unicode.In(v, unicode.Han) {
			sum++
		}
	}
	fmt.Println(sum)

	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	fmt.Println(s3)

	m := make(map[string]int, 5)
	for _, v := range s3 {
		//if m[v] == 0 {
		//	m[v] =1
		//}	else {
		//	m[v] ++
		//}
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	fmt.Println(m)

	//回文判断
	//字符串从左往右读和从右往左读是一样的，就是回文
	//黄山落叶松叶落山黄
	s4 := "黄山落叶松叶落山黄"
	s5 := make([]string, len(s4))
	for i, v := range s4 {
		s5[len(s4)-i-1] = string(v)
		//fmt.Println(i, string(v))
		fmt.Println(s5)
	}
	var s6 string
	s6 = strings.Join(s5, "")
	fmt.Println(s6)
	fmt.Println(s6 == s4)

	runes := make([]rune, 0, len(s4))
	for _, rune := range s4 {
		runes = append(runes, rune)
	}
	fmt.Println("rune[] :", runes)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return
		}
	}
	println("回文")
}
