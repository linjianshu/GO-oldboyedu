package main

import (
	"fmt"
	"os"
)

/*
函数版学生管理信息系统
写一个系统能够查看、新增、删除学生
*/

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func main() {
	for {
		fmt.Println("欢迎！")
		fmt.Println(
			`	1.查看所有
	2.新增学生
	3.删除学生`,
		)
		fmt.Println("请输入功能项")
		var d string
		fmt.Scanln(&d)
		switch d {
		case "1":
			showAllStudent1()
		case "2":
			addStudent()
		case "3":
			deleteStudent()
		case "4":
			os.Exit(1) //退出
		default:
			fmt.Println("不符合要求的输入格式")
		}
	}

}

func deleteStudent() {
	fmt.Println("请输入你删除学生的学号")
	var id int64
	fmt.Scanln(&id)
	delete(studentMap, id)
}

func addStudent() {
	var name string
	var id int64
	fmt.Println("请输入学号")
	fmt.Scanln(&id)
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	s := newStudent(id, name)
	studentMap[id] = s
}

func showAllStudent1() {
	for _, v := range studentMap {
		fmt.Printf("学号：%d  姓名：%s\n", (*v).id, (*v).name)
	}
}

var studentMap = map[int64]*student{}
