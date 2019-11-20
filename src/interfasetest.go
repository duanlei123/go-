package main

import "fmt"

// 接口 ： 描述了一系列方法的集合

type Humaner interface {
	// 方法 - 只有声明，没有实现 - 由自定义类型实现
	Sayhi()
}
// 接口继承
type Personer interface {
	Humaner //匿名字段
	sing(lrc string)
}

// 接口转换: 超集 可以转换为子集， 反之不成立.
// 空接口: 没有方法的接口，万能类型，可以保持任意类型的值
func testIner(){
	var i interface{} = 2
	fmt.Println("i = ", i)
}

// 应用场景
func  xxx(arg ...interface{}){
	// 这样就可以接收任意类型参数, [0, 多个]
}

type Student3 struct {
	name string
	id int
}
func (student *Student3) Sayhi(){
	fmt.Printf("Student[%s,%d] sayhi()\n", student.name, student.id)
}
func (student *Student3) sing(lrc string){
	fmt.Printf("Student在唱着：%s", lrc)
}

func main19() {
	var per Personer

	student8 := &Student3{
		name: "mike",
		id:   888,
	}
	per = student8
	per.Sayhi()
	per.sing("黑格")


}



//接口命名习惯以 er 结尾
//接口只有方法声明, 没有实现, 没有数据字段
// 接口可以匿名嵌入到其他接口， 或嵌入到结构中

// 接口是用来定义行为的类型, 这些被定义的行为不由接口去实现， 而是有用户定义的类型实现, 一个实现了这些方法的具体类型是这个接口类型的实例

type Student2 struct {
	name string
	id int
}

func (student *Student2) Sayhi(){
	fmt.Printf("Student[%s,%d] sayhi()\n", student.name, student.id)
}

type Teacher struct {
	addr string
	group string
}

func (teacher *Teacher) Sayhi()  {
	fmt.Printf("Teacher[%s,%s] sayhi()\n", teacher.addr, teacher.group)
}

type Mystr string

func (mystr *Mystr) Sayhi()  {
	fmt.Printf("mystr 实现了 sayhi()\n")
}

// 定义一个普通函数, 接口为形式参数
func WhoSayhi(i Humaner){
	i.Sayhi()
}

func main18() {
	student := &Student2{
		name: "mike",
		id:   666,
	}
	teacher := &Teacher{
		addr:  "北京",
		group: "go" ,
	}
	var mystr Mystr ="holle mike"

	WhoSayhi(student)
	WhoSayhi(teacher)
	WhoSayhi(&mystr)

	// 创建一个切片
	x := make([]Humaner, 3)
	x[0] = student
	x[1] = teacher
	x[2] = &mystr

	for _, i := range x{
		i.Sayhi()
	}
}


func main15() {
	//定义接口类型变量
	var i Humaner
	// 只要实现了此接口方法的类型, 那么这个类型的变量(接受者类型)就可以 i 赋值 -
	student := &Student2{
		name: "mike",
		id:   666,
	}

	i = student
	i.Sayhi()

	teacher := &Teacher{
		addr:  "北京",
		group: "go" ,
	}

	i = teacher
	i.Sayhi()

	var mystr Mystr ="holle mike"
	i = &mystr
	i.Sayhi()

	// 同一个接口， 不同的表现形式, (多态)


}


