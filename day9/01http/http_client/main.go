package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//net/http client
func main() {
	f1()
	f2()
	f3()
}

//构造一个client
var client = http.Client{Transport: &http.Transport{
	DisableKeepAlives: false}}

func f3() {
	//构造参数
	urlparams := url.Values{}
	urlparams.Add("name", "ljs")
	urlparams.Add("score", "100")
	//构造头部
	urlParse, _ := url.ParseRequestURI("http://127.0.0.1:9090/xxx")
	//拼接url
	urlParse.RawQuery = urlparams.Encode()
	//构造一个请求 request
	request, _ := http.NewRequest("get", urlParse.String(), nil)
	response, _ := client.Do(request)
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)
}

func f2() {
	//构造请求
	uri, _ := url.ParseRequestURI("http://127.0.0.1:9090/xxx")
	data := url.Values{}
	data.Set("name", "林の树")
	data.Set("age", "18")
	urlStr := data.Encode()
	fmt.Println(urlStr)
	uri.RawQuery = urlStr
	fmt.Println(uri)
	request, err := http.NewRequest("Get", uri.String(), nil)
	if err != nil {
		fmt.Println("request failed ,error : ", err)
		return
	}
	//发请求
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close() //一定要记得关闭resp.body
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))

}

func f1() {
	response, err := http.Get("http://127.0.0.1:9090/xxx?name=ljs&age=18")
	if err != nil {
		fmt.Println("get failed , error : ", err)
		return
	}
	//从response中吧服务端返回的数据读出来
	read, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read failed , error : ", err)
		return
	}

	fmt.Println(string(read))
}
