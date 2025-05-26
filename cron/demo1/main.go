package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// 创建cron对象，用于管理定时任务。
	c := cron.New()

	// 调用cron对象的AddFunc()方法向管理器中添加定时任务。
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	// 调用c.Start()启动定时循环。
	c.Start()

	// c.Start()启动一个新的 goroutine 做循环检测，在这里加入time.Sleep(time.Second * 5)防止主 goroutine 退出。
	time.Sleep(time.Second * 5)
}
