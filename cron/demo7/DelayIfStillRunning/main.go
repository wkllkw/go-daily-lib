package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

type delayJob struct {
	count int
}

func (d *delayJob) Run() {
	time.Sleep(2 * time.Second)

	d.count++

	log.Printf("%d:hello world\n", d.count)
}

func main() {
	c := cron.New()

	// DelayIfStillRunning 触发时，如果上一次任务还未执行完成（耗时太长），则等待上一次任务完成之后再执行；
	c.AddJob("@every 1s", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&delayJob{}))

	c.Start()

	time.Sleep(10 * time.Second)
}
