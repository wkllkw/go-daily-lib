package main

import (
	"context"
	"fmt"
)

func main() {
	// gen 函数在单独的goroutine中生成整数
	// 并将它们发送到返回的通道
	// gen的调用者需要在消费完生成的整数后取消context
	// 以避免gen启动的内部goroutine泄漏
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 返回以避免goroutine泄漏
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 在我们消费完整数后取消

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
