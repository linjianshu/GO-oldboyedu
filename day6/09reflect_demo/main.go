package main

import (
	"fmt"
	"reflect"
)

func main() {
	reflectType(int64(8))
	reflectType(int32(8))
	reflectType("hello world")
	reflectType(map[interface{}]interface{}{1: "hello", "hello world": []bool{true}})
	reflectType(cat{name: "tomcat"})

	reflectValue(int64(8))

	b := int64(20)
	//reflectSetValue1(b)  //这样不行会引发panic错误
	reflectSetValue2(&b)
	fmt.Println(b)

	var a *int
	fmt.Println(reflect.ValueOf(a).IsNil())
	fmt.Println(reflect.ValueOf(a).IsValid())
	c := cat{name: "tomcat"}
	fmt.Println(reflect.ValueOf(c).FieldByName("name"))
	fmt.Println(reflect.ValueOf(c).MethodByName("name").IsValid())

	m := map[string]int{"娜扎": 1}
	fmt.Println(reflect.ValueOf(m).MapIndex(reflect.ValueOf("娜扎")).IsValid())

}
func reflectType(a interface{}) {
	v := reflect.TypeOf(a)
	fmt.Printf("%T %v\n", v, v)
	fmt.Printf("type %v  kind %v \n", v.Name(), v.Kind())
}
func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	kind := v.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	}

}
func reflectSetValue1(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Int64 {
		value.SetInt(200) //修改的是副本 reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Elem().Kind() == reflect.Int64 {
		value.Elem().SetInt(200)
	}
}

type cat struct {
	name string
}
