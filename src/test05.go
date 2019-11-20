package main

import "fmt"

//1.输出9*9口诀。
//2.程序分析：分行与列考虑，共9行9列，i控制行，j控制列。

func main10() {

	for i := 1; i < 10; i++{
		for j := 1; j < i+1 ; j++  {
			result := i * j
			fmt.Printf("%d*%d=%d ", j, i, result)
		}
		fmt.Println()
	}
}
