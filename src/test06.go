package main

import "fmt"

//1.题目：要求输出国际象棋棋盘。
//2.程序分析：用i控制行，j来控制列，根据i+j的和的变化来控制输出黑方格，还是白方格。
//3.程序源代码：
func main11() {
	for i := 0; i < 8; i++{
		for j := 0; j < 8; j++{
			if (i+j)%2 == 0 {
				fmt.Printf("%c%c", 219, 219)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}


	for i := 1; i < 11; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%c%c", 219, 219)
		}
		fmt.Println()
	}
}
