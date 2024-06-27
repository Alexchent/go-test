package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds())
	spec := fmt.Sprintf("0 44 */%d * * *", 1)
	fmt.Println(spec)
	c.AddFunc(spec, func() {
		//Nu(serverCtx, hour)
		fmt.Println(time.Now().String())
	})
	c.Start()

	ch := make(chan int)

	<-ch

	ch <- 1
}
