package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func loadIni(v1 interface{}, v MysqlConf) {
	typeof := reflect.TypeOf(v)
	for i := 0; i < typeof.NumField(); i++ {
		temp := fmt.Sprint(typeof.Field(i).Tag.Get("ini"))
		i2 := m[temp]
		value := reflect.ValueOf(v1)
		fmt.Println(value)
		value1 := reflect.ValueOf(v)
		valueType := value1.Field(i).Kind()
		fmt.Println(valueType)
		if valueType == reflect.Int {
			//将interface转成int
			value.Elem().SetInt(i2.(int64))
		} else if valueType == reflect.String {
			value.Elem().SetString(i2.(string))
		}
	}
}

//ini文件解析器
func main() {
	var conf MysqlConf
	f1()
	loadIni(&conf, conf)
	fmt.Println(conf.Address)
	fmt.Println(conf.Port)
	fmt.Println(conf.UserName)
	fmt.Println(conf.Password)
}

type MysqlConf struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

var m = make(map[string]interface{}, 20)

func f1() {
	readFile, err := ioutil.ReadFile("./code.oldboyedu.com/day6/11ini_demo/conf.ini")
	if err != nil {
		return
	}
	split := strings.Split(string(readFile), "\r\n")
	for _, v := range split {
		if !strings.Contains(v, "=") {
			continue
		}
		i := strings.Split(v, "=")
		m[strings.Trim(i[0], " ")] = strings.Trim(i[1], " ")
	}
	fmt.Println(m)
}
