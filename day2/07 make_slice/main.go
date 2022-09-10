package main

import "fmt"

func main() {
	//make函数创造切片
	s1 := make([]int, 3, 10)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d,s1==nil?:%v\n", s1, len(s1), cap(s1), s1 == nil)

	var s2 []int
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d,s2==nil?:%v\n", s2, len(s2), cap(s2), s2 == nil)

	s3 := []int{}
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d,s3==nil?:%v\n", s3, len(s3), cap(s3), s3 == nil)

	s4 := make([]int, 0)
	fmt.Printf("s4=%v,len(s4)=%d,cap(s4)=%d,s4==nil?:%v\n", s4, len(s4), cap(s4), s4 == nil)

	//切片的赋值
	s5 := []int{1, 3, 5, 7}
	s6 := s5 //s5 和 s6都指向了同一个底层数组
	fmt.Println(s5, s6)
	s5[0] = 100
	fmt.Println(s5, s6)

	//切片的遍历
	//1.索引遍历
	for i := 0; i < len(s5); i++ {
		fmt.Printf("%d ", s5[i])
	}
	fmt.Println()
	//2.forrange遍历
	for _, v := range s5 {
		fmt.Printf("%d ", v)
	}
}
