package main

import (
	"encoding/json"
	"fmt"
)

// go 语言处理json数据格式

// 1, 如何生成json数据
//	1.1 通过结果体生成
//	1.2 通过map生成

// 2, 如何解析json数据

//{
//	"company":"itcast",
//	"subjects":[
//		"GO",
//		"C++",
//		"Python",
//		"Test"
//
//	],
//	"isok":true,
//	"price":666.666
//}

// 要生成json 成员首字母必须大写
//type IT struct {
//	Company string
//	Subjects [] string
//	Isok bool
//	Price float64
//}
type IT struct {
	Company string `json:"company"`
	Subjects [] string `json:"subjects"`
	Isok bool `json:"isok"` // 值可做string格式化到json
	Price float64 `json:"price"`
	//Name string `json:"-"` // 字段不会被格式化为json
}

// 通过map格式化为json


func main25() {
	// json解析 到结构体
	jsonBuf := `
		{
		 "company": "itcast",
		 "subjects": [
		  "GO",
		  "C++",
		  "Python",
		  "Test"
		 ],
		 "isok": true,
		 "price": 666.666
		}
	`
	var tmp IT
	unmarshal := json.Unmarshal([]byte(jsonBuf), &tmp)
	if unmarshal != nil{
		fmt.Println("error = ", unmarshal)
		return
	}
	fmt.Println(tmp)

	// json 解析到map
	// 创建一个map
	mapss := make(map[string]interface{}, 4)
	unmarshal1 := json.Unmarshal([]byte(jsonBuf), &mapss)
	if unmarshal1 != nil{
		fmt.Println("error = ", unmarshal1)
		return
	}
	fmt.Println(mapss)

	// 类型断言
	for key, value := range mapss{
		//fmt.Println(key,value)
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s] ====>值的类型为string, value = %s\n", key, data)
		case bool:
			fmt.Printf("map[%s] ====>值的类型为bool, value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s] ====>值的类型为float64, value = %f\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s] ====>值的类型为切片, value = %v\n", key, data)
		}
	}



	// 定义一个结构体变量 同时初始化
	//it := IT{"itcast",[]string{"GO","C++","Python","Test"}, true, 666.666, "好的"}
	//bytes, err := json.Marshal(it)
	//bytes, err := json.MarshalIndent(it, "", " ")
	//if err != nil{
	//	fmt.Println("error = ", err)
	//	return
	//}
	//fmt.Println(string(bytes))

	// 创建一个map
	//maps := make(map[string]interface{}, 4)
	//maps["company"] = "itcast"
	//maps["subjects"] = []string{"GO","C++","Python","Test"}
	//maps["isok"] = true
	//maps["price"] = 666.666

	//marshal, err := json.Marshal(maps)
	//indent, err := json.MarshalIndent(maps, "", " ")
	//if err != nil{
	//	fmt.Println("err = ", err)
	//	return
	//}
	//fmt.Println(string(indent))
}
