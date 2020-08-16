package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds())
	go try2(c)
	i := 0
	//c := cron.New()
	spec := "*/1 * * * * *"
	id, err := c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
		fmt.Println(time.Now())
	})
	fmt.Println("id-->", id)
	fmt.Println("-->", err)
	c.Start()
	go func() {
		time.Sleep(time.Second * 10)
		c.Remove(2)
	}()
	select {}
}
func try2(c *cron.Cron) {
	i := 0
	//c := cron.New()
	spec := "*/1 * * * * *"
	id, err := c.AddFunc(spec, func() {
		i++
		fmt.Println("try2:", i)
		fmt.Println(time.Now())
	})
	fmt.Println("try2-->", id)
	fmt.Println("try2-->", err)
	c.Start()

	select {}

}
