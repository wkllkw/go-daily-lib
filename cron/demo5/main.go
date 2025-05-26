package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	// 调用AddJob将GreetingJob对象添加到定时管理器中
	c.AddJob("@every 1s", GreetingJob{"dj"})

	c.Start()

	time.Sleep(time.Second * 5)
}

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Printf("Hello, %s!\n", g.Name)
}
