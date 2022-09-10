package main

import (
	"fmt"
	"strings"
)

func main() {

	//if 条件判断
	age := 19
	if age > 18 {
		fmt.Println("性感荷官在线发牌")
	} else {
		fmt.Println("好好学习，以后赌博")
	}

	if age >= 35 && age < 80 {
		fmt.Println("人到中年，不得不服")
	} else if age > 18 {
		fmt.Println("年轻力壮，不怕困难")
	} else {
		fmt.Println("好好学习，少吃点苦")
	}

	if name := "linjianshu"; strings.Contains(name, "lin") {
		fmt.Println("确实确实")
	} else {
		fmt.Println("不敢不敢")
	}

}
