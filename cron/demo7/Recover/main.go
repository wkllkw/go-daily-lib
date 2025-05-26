package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type panicJob struct {
	count int
}

func (p *panicJob) Run() {
	p.count++
	
	if p.count == 1 {
		panic("panic！！！！！！！！！！！！！！！！！")
	}

	fmt.Println("hello world")
}

func main() {
	c := cron.New()

	// Recover 捕获内部Job产生的 panic；
	c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))

	c.Start()

	time.Sleep(5 * time.Second)
}
