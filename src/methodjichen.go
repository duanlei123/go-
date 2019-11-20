package main

import "fmt"

type Person1 struct {
	name string
	age int
}

func (p *Person1) SetInfo(name string, age int)  {
	p.name = name
	p.age = age
}

func (p Person1) PrintInfo(){
	fmt.Printf("name = %s, age = %d\n", p.name, p.age)
}

// 有一个学生 ，继承了person1字段， 成员和方法都继承了
type Student1 struct {
	Person1
	id int
	addr string
}
// 方法重写
func (student *Student1) PrintInfo()  {
	fmt.Printf("name = %s, age = %d, id = %d, addr = %s\n", student.name, student.age, student.id, student.addr)
}

func main14() {
	 student := Student1{Person1{"xxx",18},001, "北京"}
	 student.PrintInfo()
}
