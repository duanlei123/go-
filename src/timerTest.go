package main

import (
	"fmt"
	"time"
)

// 定时器 timer

func main() {
	//创建一个定时器，设置时间为 2s. 2s后向timer通道写内容(当前时间)
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	<-timer.C //读取数据
	fmt.Println("时间到")

	<- time.After(time.Second)
}
