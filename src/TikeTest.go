package main

import (
	"fmt"
	"time"
)

// 定时器 Tick
func main() {
	tick := time.NewTicker(time.Second)

	i := 0
	for {
		<-tick.C
		i++
		fmt.Println("i = ", i)
		if i == 5{
			tick.Stop()
			break
		}
	}
}
