package main

import "fmt"

//1.题目：有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
//
//2.程序分析：可填在百位、十位、个位的数字都是1、2、3、4。组成所有的排列后再去掉不满足条件的排列。
//
//3.程序源代码：

func main07(){
	number := 0
	for i := 1; i < 5; i++{
		for j := 1; j < 5 ; j++  {
			for k := 1; k < 5; k++ {
				if i != k && i != j && j != k{
					fmt.Println("这个数是：", i,j,k)
					number ++
				}
			}
		}
	}
	fmt.Println("共有生成三位数：", number)
}
