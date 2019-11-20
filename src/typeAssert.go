package main

import "fmt"

// 类型断言

// 通过if来实现
type User struct {
	name string
	id int
}


func main20() {
	slice1 := make([]interface{},3)
	slice1[0] = 1
	slice1[1] = "hello go"
	slice1[2] = User{"mayun", 45}

	// 类型查询- 类型断言

	for index, data := range slice1{
		fmt.Printf("index = %d, data = %v\n", index, data)
		if value, ok := data.(int);ok == true{
			fmt.Printf("slice[%d] 类型为int, 内容为: %v", index, value)
		}else if value,ok := data.(string); ok == true{
			fmt.Printf("slice[%d] 类型为string, 内容为: %v", index, value)
		}else if value,ok := data.(User);ok == true {
			fmt.Printf("slice[%d] 类型为User, 内容为: %v", index, value)
		}
	}
}
