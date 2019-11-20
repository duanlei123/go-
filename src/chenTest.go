package main

import (
	"fmt"
	"time"
)

func main31() {

	// 无缓存通道
	ch := make(chan int, 3)
	//fmt.Printf("len(ch) = %d, cap(ch) = %d", len(ch), cap(ch))

	// 创建协程
	go func() {
		for i := 0; i < 3 ; i++  {
			//fmt.Println("子协程： i = ", i)
			ch <- i // 将i写入通道
			//fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
			fmt.Println("子协程[%d]", len(ch))
		}
		close(ch) //关闭ch
	}()

	// 延时
	time.Sleep(2 * time.Second)

	for{
		if num, ok := <- ch; ok == true{ // 为true 时 管道没有关闭。
			fmt.Println("num = ", num)
		}else {
			break
		}
	}

	//// 创建通道
	//ch := make(chan string)
	//go func(){
	//	defer fmt.Println("子线程调用完毕")
	//	for i :=0; i < 2; i++{
	//		fmt.Println("子线程 i = ", i)
	//		time.Sleep(time.Second)
	//	}
	//	ch <- "我是子完了，你也完吧"
	//}()
	//
	//str := <- ch // 没有数据会阻塞
	//fmt.Println("str = ", str)
}
