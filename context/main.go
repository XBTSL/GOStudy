package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var syncG sync.WaitGroup

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	syncG.Add(1)
	go a(ctx)
	context.Canceled = errors.New("Mq 故障......")
	ctx, cancle := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("停止消费", ctx.Err(), "consumer", i, "\n")
				default:
					fmt.Println("消费topic:order中的消息")
					time.Sleep(time.Second)
				}
			}
		}(i)
	}
	go func() {
		time.Sleep(time.Second * 1)
		// 故障，通知消费者不在消费
		cancle()
	}()

	// 暂时别让程序停止
	time.Sleep(time.Second * 2)
	syncG.Wait()

}
func a(ctx context.Context) {
	fmt.Println("a 方法调用的来源", ctx.Value("key"))
	ctx = context.WithValue(ctx, "key", "valuesaa")
	go b(ctx)
}

func b(ctx context.Context) {
	fmt.Println("b 方法调用来源", ctx.Value("key"))
	syncG.Done()
}
