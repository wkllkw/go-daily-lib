package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用 sync.WaitGroup 来等待 goroutine 结束
var wg sync.WaitGroup

// worker 函数模拟一个数据库连接工作
func worker(ctx context.Context) {
LOOP: // 使用标签标识循环，方便后面跳出
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 模拟每次连接数据库耗时10毫秒

		select {
		case <-ctx.Done(): // 监听context取消信号
			break LOOP // 收到信号后跳出循环
		default:
			// 没有收到取消信号则继续工作
		}
	}
	fmt.Println("worker done!")
	wg.Done() // 通知WaitGroup当前goroutine已完成
}

func main() {
	// 创建一个50毫秒超时的context
	// WithTimeout 实际上返回的是 WithDeadline(now + timeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	// 启动worker goroutine
	wg.Add(1) // WaitGroup计数器+1
	go worker(ctx)

	// 主程序休眠5秒（远大于worker的超时时间）
	time.Sleep(time.Second * 5)

	// 显式调用cancel函数（良好实践，即使可能已经超时）
	cancel()

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("over")
}
