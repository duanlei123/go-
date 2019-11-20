package main

import "fmt"

// recover() 函数, panic 出现 程序就回崩溃，
// recover() 就是用来拦截运行时异常的一个内置函数， 它可以是当前的程序从运行时panic的状态中恢复并重新获取流程控制权。

// 注意, 注意 recover 只有在defer调用的函数中有效.

func main22() {
	// 数组越界导致panic
	testa()
	testb(20)
	testc()
}

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaa")
}

func testb(x int) {
	fmt.Println("bbbbbbbbbbbbbbbbbbbbb")
	// 设置 recover
	defer func() {
		//recover() // 可以打印错误信息的
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 111
}

func testc() {
	fmt.Println("ccccccccccccccccccccc")
}

