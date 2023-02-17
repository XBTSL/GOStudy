package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func requestByParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("name", "xiongbin")
	params.Add("age", "18")
	params.Encode()
	request.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do((request))
	if err != nil {
		panic(err)
	}
	printBody(r)
}

func requestByHead() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("user-agent", "chrome")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	printBody(r)
}

func main() {
	//如何查询请求参数
	//如何定制请求头，比如修改user-agent
	requestByParams()
	requestByHead()
}
