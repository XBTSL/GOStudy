package main

import "fmt"

func main() {
	fmt.Println("demo1:", demo1())
	fmt.Println("demo2:", demo2())
}

func demo1() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func demo2() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}
