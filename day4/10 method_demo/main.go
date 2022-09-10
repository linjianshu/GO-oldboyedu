package main

import "fmt"

//标识符：变量名、函数名、类型名、方法铭
//go语言中如果标识符首字母是大写的，就表示对外部包可见（暴露的，公有的）

//Dog 这是一个狗的结构体注释
type Dog struct {
	name string
}

func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//方法是作用于特定类型的函数
//接受者表示的是调用该方法的具体类型变量，多用类型变量首字母小写表示
func (d Dog) wangwang() {
	fmt.Println(d.name + "汪汪汪")
}

//使用值接收者：传拷贝进去
func (p person) guonian() {
	p.age++
}

//操作指针 指针接收者：传地址进去
func (p *person) guonian1() {
	(*p).age++
}
func main() {
	newDog("jwt").wangwang()

	p := newPerson("ljs", 18)
	fmt.Println(p.age)
	p.guonian()
	fmt.Println(p.age)

	p1 := newPerson("jwt", 19)
	fmt.Println(p1.age)
	p1.guonian1()
	fmt.Println(p1.age)

}
