package main

import (
	"calc"
	"fmt"
	"math/rand"
	"time"
)
//同一个目录调用别的文件是函数不需要包名引用

//func init() {
//	fmt.Println("main包中init()函数执行")
//}

func main1()  {
	//fmt.Println("我的唯一入口")
	//test01()
	i := calc.Add(2 , 3)
	fmt.Println(i)
	fmt.Println(calc.Minus(5,1))

	fmt.Println("----------------------------")

	// 指针 pointer : 代表某个内存地址
	var  n int  = 10
	fmt.Printf("&a = %p\n", &n)

	var p *int
	p = &n
	fmt.Printf("p = %v\n", p)

	*p = 666
	fmt.Printf("*p = %v, n = %v\n", *p, n)

	var m *int
	m = new(int)
	*m = 666
	fmt.Println("*m = ", *m)

	fmt.Println("================================")

	a, b := 10 , 20
	// 通过一个函数交换a 和 b 的内容
	//swap(a,b)
	//fmt.Printf("main a = %d, b = %d\n", a, b)

	swapL(&a,&b)
	fmt.Printf("main* a = %d, b = %d\n", a, b)

	// 定义一个数组
	var id [50]int

	for i :=0; i<len(id); i++{
		id[i] = i+1
	}
	fmt.Println(id)
	fmt.Printf("len(id) = %d\n", len(id)) //注意数组元素的个数必须是常量、

	//数组初始化
	// 全部初始化
	var array [5]int = [5]int{1,2,3,4,5}
	fmt.Println(array)
	//自动推到类型
	array1 := [5]int{1,2,3,4,5}
	fmt.Println(array1)

	random()
}

func swap(a, b int)  {
	a, b = b, a
	fmt.Printf("swap a = %d, b = %d\n", a, b)
}

func swapL(a , b *int)  {
	*a, *b = *b, *a
	fmt.Printf("swap* a = %d, b = %d\n", *a, *b)
}
//随机数
func random(){
	//设置种子，只需一次
	rand.Seed(time.Now().UnixNano()) //如果种子一样，每次运行程序产生的随机数都一样，time.Now().UnixNano()当前时间
	var a [10]int
	for i := 0; i< len(a) ;i++  {
		//i := rand.Int()
		//fmt.Println("rand = ", i)
		//intn := rand.Intn(100) //限制在100
		a[i] = rand.Intn(100)
		fmt.Printf("%d,", a[i])
	}
	fmt.Printf("\n")
	// 冒泡排序
	//bubblesort(a)
}
// 冒泡排序(升须), 数组做参数 是属于值传递,
func  bubblesort(n []int) {
	for i:=0; i<len(n)-1 ; i++  {
		for j:=0;j<len(n)-1-i ; j++  {
			if(n[j] > n[j+1]){
				n[j], n[j+1] = n[j+1],n[j]
			}
		}
	}
	fmt.Println("排序后：", n)
}
//切片
func main02()  {
	// 声明切片,
	ss := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(ss), cap(ss))

	// 添加元素
	ss = append(ss, 11)
	fmt.Printf("len = %d, cap = %d\n", len(ss), cap(ss))

	// 借助make函数创建切片
	sss := make([]int, 5, 10)
	fmt.Printf("len = %d, cap = %d\n", len(sss), cap(sss))

	ssss := make([]int, 5)
	fmt.Printf("len = %d, cap = %d\n", len(ssss), cap(ssss))


	a := []int{0,1,2,3,4,5,6,7,8,9}
	s := a[2:5]
	fmt.Println(s)
	s2 := s[2:7]
	fmt.Println(s2)
	s2[2] = 777
	fmt.Println(s2)
	fmt.Println(a)

	sd := make([]int, 0 , 1)
	oldCap := cap(sd)
	for i := 0; i< 8;i++{
		sd = append(sd, i)
		if newCap := cap(sd); oldCap < newCap{
			fmt.Printf("cap: %d =====> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}

}
func main03(){
	srcSlice := []int{1,2}
	dstSlice := []int{6,6,6,6,6}
	copy(dstSlice,srcSlice)
	fmt.Println("dst:",dstSlice)

	// 切片做函数参数传递, 切片为引用类型, 直接为引用传递
	var n int = 10
	s := make([]int, n)
	InitData(s) // 初始化数组
	fmt.Println("排序前：", s)
	bubblesort(s)
	fmt.Println("排序后：", s)
}
func InitData(s []int)  {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(s) ;i++  {
		s[i] = rand.Intn(100)
	}
}
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	//panic("触发异常")
}

// 猜数字小游戏
func main04() {
	// 随机数
	var randNum int
	// 随机产生一个4位随机数
	CreatNum(&randNum)
	//fmt.Println(randNum)
	randSlice := make([]int, 4)
	// 保存这个4位数的每一位
	GetNum(randSlice, randNum)
	fmt.Println("randSlic = ", randSlice)
	OnGano(randSlice) //游戏

}
func OnGano(randSlice []int){
	var number int
	keySlice := make([]int, 4)
	for{
		for{
			fmt.Println("请输入一个4位数:")
			fmt.Scan(&number)
			if 999 < number && number < 10000{
				break
			}
			fmt.Println("输入的数不符合要求！请重新输入")
		}
		GetNum(keySlice, number)
		n := 0
		for i:=0; i<4 ; i++  {
			if keySlice[i] > randSlice[i]{
				fmt.Printf("第%d位大了\n", i+1)
			}else if keySlice[i] < randSlice[i]{
				fmt.Printf("第%d位小了\n", i+1)
			}else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}
		if n == 4{
			fmt.Println("全部猜对了")
			break
		}
	}
}
func GetNum(s []int, num int)  {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}
func CreatNum(p *int)  {
	rand.Seed(time.Now().UnixNano())
	var num int
	for{
		num = rand.Intn(10000)
		if num >= 1000{
			break
		}
	}
	//fmt.Println(num)
	*p = num
}

// map 一个无序的 key-value 对的集合
func main05()  {
	var m1  map[int]string
	fmt.Println(m1)

	m2 := make(map[int]string)
	fmt.Println(m2)
	fmt.Println(len(m2))

	m3 := make(map[int]string, 2)
	m3[1] = "mike"
	m3[2] = "go"
	m3[3] = "c++"
	fmt.Println(m3)
	fmt.Println(len(m3))

	m4 := map[int]string{1:"001",2:"002",3:"003",4:"004"}
	fmt.Println(m4)

	//遍历
	for key, value:= range m4{
		fmt.Println(key,value)
	}

	//判断key是否存在
	value , ok := m4[5]
	if ok {
		fmt.Println(value)
	}else {
		fmt.Println("key 不存在")
	}

	// 删除可以
	delete(m4, 1)
	fmt.Println(m4)

	// 做函数参数 ， 引用传递
	test(m4) // 在函数内部删除key为 4
	fmt.Println(m4)
}
func test(m map[int]string){
	delete(m, 4)
}
//结构体
// 初始化结构体
//type  Student struct {
//	id int
//	name string
//	sex byte
//	age int
//	addree string
//}
//func main06(){
//	//结构体是一种聚合的数据类型, 他是由一系列具有相同类型或不同类型的数据结构集合， 每个数据称为结构体的成员
//	// 顺序初始化, 每个成员必须初始化
//	var p1 *Student = &Student{1,"duanlei",'m',18,"北京"}
//	fmt.Println(*p1)
//
//	// 指定成员初始化， 没有初始化的成员自动赋值为0
//	p2 := &Student{name:"张飞", addree:"天津"}
//	fmt.Printf("p2 type is: %t\n", *p2)
//	fmt.Println(*p2)
//
//	// 成员使用
//	var s Student
//	s.name="李四"
//	s.id=3
//	s.age=19
//	s.addree="啥关系"
//	s.sex='m'
//	fmt.Println(s)
//
//	// 指针有合法指向, 才操作成员
//	var s1 Student
//	var p *Student
//	p = &s1
//
//	p.id = 5
//	(*p1).name="张飞"
//	p.addree = "世纪东方"
//	p.age = 14
//	p.sex = 's'
//	fmt.Println(*p)
//
//	p4 := new(Student)
//	p4.id = 6
//	fmt.Println(p4)
//
//	// 函数参数
//	test02(s) // 值传递
//	fmt.Println(s)
//
//	test03(&s)
//	fmt.Println(s)
//
//}
//func test03(s *Student){
//	s.id = 666
//	fmt.Println("引用传递",s)
//}
//func test02(s Student){
//	s.id = 666
//	fmt.Println("值传递",s)
//}


