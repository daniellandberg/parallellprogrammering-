package main

import (
	"fmt"
	"time"
)

func main() {
	eat := time.Tick(10 * time.Second)
	work := time.Tick(30 * time.Second)
	sleep := time.Tick(60 * time.Second)
	t := time.Now()
	for {
		select {
		case <-eat:
			fmt.Println("the time is ", t.Format("15 4 5"), ":", "Time to eat")
		case <-work:
			fmt.Println("the time is ", t.Format("15 4 5"), ":", "time to work")
		case <-sleep:
			fmt.Println("the time is ", t.Format("15 4 5"), ":", "time to sleep")

		}
	}}