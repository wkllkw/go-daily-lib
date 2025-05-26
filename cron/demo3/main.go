package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {

	c := cron.New()

	c.AddFunc("@daily", func() {
		Hello()
		fmt.Println("tick everyday")
	})

	c.AddFunc("@weekly", func() {
		Hello()
		fmt.Println("tick weekday")
	})

	c.AddFunc("@monthly", func() {
		Hello()
		fmt.Println("tick monthday")
	})

	c.AddFunc("@yearly", func() {
		Hello()
		fmt.Println("tick yearday")
	})
	
	c.Start()

	for {
		time.Sleep(time.Second)
	}
}

func Hello() {
	fmt.Println("hello world")
}
