package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//net / http server
func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("0.0.0.0:9090", nil)
}

func f1(w http.ResponseWriter, r *http.Request) {

	//file1, err := ioutil.ReadFile("./mydear.jpg")
	//if err != nil {
	//	w.Write([]byte(fmt.Sprintf("%v",err)))
	//	return
	//}
	//
	//w.Write(file1)

	w.Header().Set("Content-Type", "image/jpg")
	imgpath := "./mydear.jpg"
	http.ServeFile(w, r, imgpath)

	file, err := ioutil.ReadFile("./index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	w.Write(file)
}
func f2(writer http.ResponseWriter, request *http.Request) {
	//对于get请求 参数都放在URL上 （query param）请求体中是没有数据的
	fmt.Println(request.URL)
	fmt.Println(request.Method)
	fmt.Println(ioutil.ReadAll(request.Body))
	writer.Write([]byte("ok!"))

	queryParam := request.URL.Query()
	fmt.Println(queryParam)
	name := queryParam.Get("name")
	fmt.Println(name)
	age := queryParam.Get("age")
	fmt.Println(age)
}
