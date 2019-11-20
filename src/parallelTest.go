package main

import (
	"fmt"
)

// goroutine : 是go并发设计的核心, 比线程小,
// 在并发编程中, 我们通常想将一个过程切分成几块，然后让每个goruitine各自负责一块工作， 当一个程序启动时，其主函数就在一个单独的goroutine中运行
// 就是main goroutine。

// 新的goroutine 用 go语句来创建

// runtime 包： 三个函数： runtime.Goched() 用于让出CPu时间片, 让出当前goroutine的执行权限, 调度器安排其他等待任务执行, 并在下一次某个时候从该位置恢复执行.
//						  runtime.Goexit() 结束当前goroutine

// chen 通道，实现多goroutine同步，保证了多goroutine并发安全

// 创建一个chen
var ch = make(chan int)

// 定义一个打印机
func Printer(str string)  {
	for _, data := range str{
		fmt.Printf("%c", data)
	}
}

func User1()  {
	Printer("hello")
	ch <- 666
	//wg1.Done() // 操作完成，减少一个计数
}

func User2()  {
	<- ch // 从管道取数据，如果通道没有数据就会阻塞
	Printer("world")
}

func main30() {
	go User1()
	go User2()
	var i = 1
	for {
		// will block here, and never go out
		i++
	}
}