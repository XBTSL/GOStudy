package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	r, err := http.Get("http://httpbin.org/get")
	defer func() {
		r.Body.Close()
	}()
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func post() {
	r, err := http.Post("http://httpbin.org/post", "", nil)
	defer func() {
		r.Body.Close()
	}()
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	defer func() {
		r.Body.Close()
	}()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func del() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	defer func() {
		r.Body.Close()
	}()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

//type student struct {
//}
//
//func (a *student) ServeHTTP(rp http.ResponseWriter, rq *http.Request) {
//	fmt.Fprintln(rp, "hello tomjack")
//}
//
//func con(rp http.ResponseWriter, rq *http.Request) {
//	fmt.Println(rp, "congratulate!!!\n")
//}

func main() {
	//temp := student{}
	//http.Handle("/", &temp)
	//http.HandleFunc("master", con)
	//http.ListenAndServe(":8080", nil)
	del()
}
