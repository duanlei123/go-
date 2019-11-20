package main

import (
	"fmt"
	"time"
)

// select 的作用： 可以监听 chan 上的数据流动
func main33() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for{
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(5* time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()
	<-quit
	fmt.Println("程序结束")
}
