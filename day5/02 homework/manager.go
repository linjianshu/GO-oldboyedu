package main

import "fmt"

//学生
type student struct {
	id    int
	name  string
	score byte
}

//学生构造函数
func newStudent(id int, name string, score byte) student {
	return student{id: id, name: name, score: score}
}

//管理员
type manager struct {
	allStudent map[string]student
}

// AddStudent 方法
func (m manager) AddStudent(s student) {
	m.allStudent[s.name] = s
	fmt.Println("添加成功！")
}

func (m manager) EditStudent(s student) {
	_, ok := m.allStudent[s.name]
	if ok {
		m.allStudent[s.name] = s
		fmt.Println("修改成功")
	} else {
		fmt.Println("无此人")
	}
}
func (m manager) DeleteStudent(name string) {
	delete(m.allStudent, name)
	fmt.Println("删除成功")
}

func (m manager) GetAllStudent() {
	for _, v := range m.allStudent {
		fmt.Printf("id:%d  name:%s  score:%d\n", v.id, v.name, v.score)
	}
}
