package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"sync/atomic"
	"time"
)

type skipJob struct {
	count int32
}

func (s *skipJob) Run() {
	atomic.AddInt32(&s.count, 1)

	log.Printf("%d:hello world\n", s.count)

	if atomic.LoadInt32(&s.count) == 1 {
		time.Sleep(2 * time.Second)
	}

}

func main() {
	c := cron.New()

	// SkipIfStillRunning 触发时，如果上一次任务还未完成，则跳过此次执行。
	c.AddJob("@every 1s", cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&skipJob{}))

	c.Start()

	time.Sleep(5 * time.Second)

}
