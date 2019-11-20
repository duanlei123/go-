package main

import "fmt"

//1.题目：输入某年某月某日，判断这一天是这一年的第几天？
//2.程序分析：以3月5日为例，应该先把前两个月的加起来，然后再加上5天即本年的第几天，
//特殊情况，闰年且输入月份大于3时需考虑多加一天。

func main()  {
	var year, mouth, day, sum, feb int
	fmt.Println("请输入年月日：")
	fmt.Scan(&year, &mouth, &day)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		feb = 29
	} else {
		feb = 28
	}
	switch mouth {
	case 1:
		sum = 0 + day
	case 2:
		sum = 31 + day
	case 3:
		sum = feb + 31 + day
	case 4:
		sum = 31 + feb + 31 + day
	case 5:
		sum = 30 + 31 + feb + 31 + day
	case 6:
		sum = 31 + 30 + 31 + feb + 31 + day
	case 7:
		sum = 30 + 31 + 30 + 31 + feb + 31 + day
	case 8:
		sum = 31 + 30 + 31 + 30 + 31 + feb + 31 + day
	case 9:
		sum = 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day
	case 10:
		sum = 30 + 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day
	case 11:
		sum = 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day
	case 12:
		sum = 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day
	}
	fmt.Printf("您输入的日期是这一年的第%d天", sum)
}