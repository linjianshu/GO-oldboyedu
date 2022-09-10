package main

import (
	json2 "encoding/json"
	"fmt"
	"reflect"
)

func main() {
	s := student{
		Name:  "xwz",
		Score: "90",
	}
	//最终要得到 {"name":"xwz","score":90}
	typeof := reflect.TypeOf(s)
	json := `{`
	for i := 0; i < typeof.NumField(); i++ {
		fmt.Println(typeof.FieldByIndex([]int{i}).Name)
		//fmt.Println(typeof.Field(i).Name)
		fmt.Println(typeof.FieldByIndex([]int{i}).Tag.Get("json"))
		json += "\"" + typeof.Field(i).Tag.Get("json") + "\"" + ":"
		structField, b := typeof.FieldByName(typeof.Field(i).Name)
		fmt.Println(structField.Name)
		fmt.Println(structField.Type)
		fmt.Println(structField.Index)
		valueof := reflect.ValueOf(s)
		fmt.Println(valueof.Field(i))
		sprint := fmt.Sprint(valueof.Field(i))
		if b {
			json += "\"" + sprint + "\"" + ","
		}
	}
	s2 := json[:len(json)-1]
	json = s2
	json += "}"
	fmt.Println(json)

	var s1 student
	json2.Unmarshal([]byte(json), &s1)
	fmt.Println(s1)
	//反序列化实例
}

type student struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}
