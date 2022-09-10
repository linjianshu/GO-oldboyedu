package main

import "fmt"

//常量
const pi = 3.1415926

//常量定义了之后不能修改
//在程序运行期间不会改变

//批量声明常量时，如果某一行没有赋值，默认就和上一行一致
const (
	pi1 = pi
	pi2
	pi3
)

const (
	i1 = iota
	i2
	i3
)

const (
	n1 = iota
	n2
	_
	n3
)

//插队
const (
	k1 = iota
	k2 = 100
	k3
	k4 = iota
	k5
)

const (
	p1, p2 = iota + 1, iota + 2
	p3, p4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func main() {
	// 不可以 pi = 12.3
	fmt.Println("pi2:", pi2)

	fmt.Println("i1:", i1)
	fmt.Println("i2:", i2)
	fmt.Println("i3:", i3)

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("k1:", k1)
	fmt.Println("k2:", k2)
	fmt.Println("k3:", k3)
	fmt.Println("k4:", k4)
	fmt.Println("k5:", k5)

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
	fmt.Println("p3:", p3)
	fmt.Println("p4:", p4)

	fmt.Println("kb:", KB)
	fmt.Println("mb:", MB)
	fmt.Println("gb:", GB)
	fmt.Println("tb:", TB)

}
