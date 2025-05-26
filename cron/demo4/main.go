package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	nyc, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(nyc))
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second at New_York")
	})

	c.AddFunc("CRON_TZ=Asia/Tokyo @every 1s", func() {
		fmt.Println("tick every 1 second at Tokyo")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
