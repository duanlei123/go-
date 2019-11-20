package main

import "fmt"

//1.题目：输入三个整数x,y,z，请把这三个数由小到大输出。
//2.程序分析：我们想办法把最小的数放到x上，
//先将x与y进行比较，如果x>y则将x与y的值进行交换，
//然后再用x与z进行比较，如果x>z则将x与z的值进行交换，这样能使x最小。

func main09() {
	var x, y, z int
	fmt.Println("请输入三个整数:")
	fmt.Scan(&x, &y, &z)

	if x < y {
		x, y = y, x
	}
	if x < z {
		x, z = z, x
	}
	if y < z {
		y, z = z, y
	}
	fmt.Println("这三个数的从大到小的顺序是：", x, y, z)
}