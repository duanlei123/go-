package main

import "fmt"

//func test01()  {
//	fmt.Println("我是test")
//}

// 匿名字段，实现继承
type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person // 匿名字段
	id int
	addr string
}
// 带有接收者的函数叫方法
func (p Person) PrintInfo(){
	fmt.Println(p)
}

// 通过一个函数, 改成员赋值
func (p *Person) SetInfo(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
}

func main13()  {

	p := &Person{"haha", 'm',2}
	p.SetInfo("heh", 'f', 18)
	p.PrintInfo()


	// 定义并初始化
	//p := Person{"duanlei",'m',18}
	//p.PrintInfo()
	//
	//var p2 Person
	//(&p2).SetInfo("youyou", 'f',22)
	//p2.PrintInfo()

	//// 初始化 Person 和 student
	//var s1 Student = Student{Person{"李四",'m',18},1,"北京"}
	//fmt.Println(s1)
	//
	////自动推到类型
	//s2 := Student{Person{"李四2",'x',19},1,"北京"}
	//fmt.Println(s2)
}


