package main

import "fmt"

func main() {
	newOptions("str1", "str2", "str3", 1, 2, 3)
	newOptionsNew(WithStrOption("str1"))
}

type OptionsNew struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// Option 先声明一个函数类型,用于传参
type Option func(option *OptionsNew)

// WithStrOption 定义具体给某个字段赋值的方法 返回一个方法 通过这个方法给结构体赋值
func WithStrOption(str string) Option {
	return func(option *OptionsNew) {
		option.strOption1 = str
	}
}

//初始化结构体
func newOptionsNew(otions ...Option) {
	options := &OptionsNew{}
	//遍历otions,得到每一个函数
	for _, fun := range otions {
		//调用函数, 在函数里,给传进去的对象赋值
		fun(options)
	}
	fmt.Printf("init options %#v\n", options)
}

type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

func newOptions(strOption1, strOption2, strOption3 string, intOption1, intOption2, intOption3 int) {
	options := Options{
		strOption1: strOption1,
		strOption2: strOption2,
		strOption3: strOption3,
		intOption1: intOption1,
		intOption2: intOption2,
		intOption3: intOption3,
	}

	fmt.Printf("init option %#v\n", options)
}
