package main

import "fmt"

// 单向chan应用

func main() {
	// 创建一个协程
	ch := make(chan int)
	//数据生产者
	go producer(ch)
	//数据消费者
	consumer(ch)
}

func consumer(ints <-chan int) {
	for num := range ints {
		fmt.Println(num)
	}
}

func producer(ints chan<- int) {
	for i := 0; i < 9 ; i++  {
		ints <- i * i
	}
	close(ints)
}
