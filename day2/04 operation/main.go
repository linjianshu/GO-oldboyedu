package main

import "fmt"

func main() {

	//运算符
	var (
		a = 5
		b = 2
	)

	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在=的右边赋值
	b++

	//关系运算符
	fmt.Println(a == b) //go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班族")
	} else {
		fmt.Println("不用上班")
	}

	if age > 60 || age < 18 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("上班族")
	}

	//not取反
	b2 := true
	fmt.Println(!b2)

	//位运算：针对的是二进制数
	//5的二进制表示 101
	//2的二进制表示 010
	//按位与
	fmt.Println(101 & 010)
	fmt.Println(5 & 2)
	//按位或
	fmt.Println(101 | 10)
	fmt.Println(5 | 2)
	//^按位异或
	fmt.Println(101 ^ 010) //这个有点奇怪 这个是109答案？？？
	fmt.Println(5 ^ 2)

	//左移右移运算 *2 和 \2
	fmt.Println(5 << 2) //101=>10100
	fmt.Println(1 << 10)
	fmt.Println(5 >> 1) //101=>10

	//注意别溢出了
	m := int8(1)
	fmt.Println(m << 10)
	fmt.Println(1<<2 + 1)

	//192.168.1.1
	//权限 文件操作会将位运算实际的应用
	//0644
	//赋值运算符，用来给变量赋值的
	var x int
	x = 10
	fmt.Println(x)
	x += 1
	fmt.Println(x)
	x -= 1
	fmt.Println(x)
	x *= 2
	fmt.Println(x)
	x /= 2
	fmt.Println(x)
	x <<= 2
	fmt.Println(x)
	x >>= 2
	fmt.Println(x)
	fmt.Printf("%b", x)
	fmt.Println()
	x &= 2
	fmt.Println(x)
	fmt.Printf("%b", x)
	x |= 2
	x <<= 2
	x >>= 2
	x ^= 2

}
