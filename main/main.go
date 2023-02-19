package main

import "fmt"

func main() {
	result := isPalindrome(121)
	fmt.Println(result)
}

func isPalindrome(x int) bool {
	temSlice := make([]int, 0)
	for x > 0 {
		tem := x % 10
		temSlice = append(temSlice, tem)
		x = x / 10
	}
	length := len(temSlice)
	i := 0
	for length-1 > i {
		if temSlice[i] == temSlice[length-i-1] {
			i++
		} else {
			return false
		}
	}
	return true
}
