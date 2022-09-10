package main

import (
	"fmt"
	"os"
)

//菜单函数
func showMenu() {
	fmt.Println("welcome!")
	fmt.Println(`1.查看所有
					 2.添加学生
					 3.修改学生
					 4.删除学生
					 5.退出`)
}
func main() {
	//学生管理系统
	//有一个物件
	//1.它保存了一些数据   -->结构体的字段
	//2.它有三个功能      -->结构体的方法
	m := manager{allStudent: map[string]student{}}
	for {
		showMenu()
		fmt.Println("输入功能项")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			m.GetAllStudent()
		case 2:
			fmt.Println("请输入学生学号，姓名，成绩")
			var id int
			var name string
			var score byte
			fmt.Scanln(&id)
			fmt.Scanln(&name)
			fmt.Scanln(&score)
			s := newStudent(id, name, score)
			m.AddStudent(s)
		case 3:
			fmt.Println("请输入学生学号，姓名，成绩")
			var id int
			var name string
			var score byte
			fmt.Scanln(&id)
			fmt.Scanln(&name)
			fmt.Scanln(&score)
			s := newStudent(id, name, score)
			m.EditStudent(s)
		case 4:
			fmt.Println("请输入学生姓名")
			var name string
			fmt.Scanln(&name)
			m.DeleteStudent(name)
		case 5:
			os.Exit(1)
		default:
			fmt.Println("格式错误")
		}

	}
}
