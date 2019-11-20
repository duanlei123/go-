package main

import "fmt"
import "errors"

// go 语言异常处理
// error 接口

// panic 当遇到不可恢复的错误状态的时候， 如数组角标越界，空指针应用等， 这些运行时错误会引起 painc 错误,
// error 接口处理方式显然不合适了， 我们不应该通过调用 panic函数来报告普通错误， 而应该把它作为报告致命错误的方式
// panic 异常发生，程序会中断,

func MyDiv(a , b int) (result int, err error){
	if b == 0 {
		err = errors.New("分母不能为0")
	}else {
		result =  a / b
	}
	return
}

func main21() {

	errorf := fmt.Errorf("%s", "this is normol err")
	fmt.Println("errorf=", errorf)

	errorf = errors.New("this is errors err")
	fmt.Println(errorf)

	//应用
	result, err := MyDiv(10, 0)
	if err != nil{
		fmt.Println("err=",err)
	}else {
		fmt.Println("result =", result)
	}

}

