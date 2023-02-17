package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	// 1. 写入
	m.Store("test", 18)
	m.Store("mo", 20)

	// 2. 读取
	age, _ := m.Load("test")
	fmt.Println(age.(int))

	// 3. 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	// 4. 删除
	m.Delete("test")
	age, ok := m.Load("test")
	fmt.Println(age, ok)

	// 5. 读取或写入
	m.LoadOrStore("mo", 100)
	age, _ = m.Load("mo")
	fmt.Println(age)
}
