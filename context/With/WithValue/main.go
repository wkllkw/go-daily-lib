package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 定义一个类型用于context的key，避免直接使用字符串可能导致的冲突
type TraceCode string

// 使用 sync.WaitGroup 来等待 goroutine 结束
var wg sync.WaitGroup

// worker 函数模拟一个工作流程，可以从context中获取跟踪码
func worker(ctx context.Context) {
	// 定义要查找的key
	key := TraceCode("TRACE_CODE")

	// 从context中获取跟踪码，并进行类型断言
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

LOOP: // 使用标签标识循环，方便后面跳出
	for {
		// 打印工作状态和跟踪码
		fmt.Printf("worker, trace code:%s\n", traceCode)

		// 模拟每次工作耗时10毫秒
		time.Sleep(time.Millisecond * 10)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	// 在context中设置跟踪码值，使用自定义类型作为key
	// 这样可以避免不同包之间的key冲突
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")

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
