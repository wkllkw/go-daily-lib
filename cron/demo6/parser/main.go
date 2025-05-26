package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	test01()

	test02()

}

func test01() {
	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	c := cron.New(cron.WithParser(parser))

	c.AddFunc("1 * * * *", func() {
		fmt.Println("every 1 second")
	})

	c.Start()

	time.Sleep(5 * time.Second)
}

func test02() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("2 * * * *", func() {
		fmt.Println("every 2 seconds")
	})

	c.Start()

	time.Sleep(5 * time.Second)
}
