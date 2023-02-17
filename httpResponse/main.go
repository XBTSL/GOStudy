package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func requsetBody(r *http.Response) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func status(r *http.Response) {
	fmt.Println(r.StatusCode) //状态码
	fmt.Println(r.Status)     //状态信息
}

func header(r *http.Response) {
	fmt.Println(r.Header.Get("Content-type"))
}

func encoding(r *http.Response) {

}

func main() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() {
		r.Body.Close()
	}()
	requsetBody(r)
	status(r)
	header(r)
}
