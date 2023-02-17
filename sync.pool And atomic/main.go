package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var caleNums int32

func createNums() interface{} {
	atomic.AddInt32(&caleNums, 1) //原子操作
	buff := make([]byte, 1024)
	return &buff
}
func main() {
	bufferPool := &sync.Pool{ //缓存池
		New: createNums,
	}

	numberWorker := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numberWorker)
	for i := 0; i < numberWorker; i++ {
		go func() {
			buffer := bufferPool.Get()
			_ = buffer.(*[]byte)
			defer wg.Done()
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	fmt.Println(caleNums)
}
