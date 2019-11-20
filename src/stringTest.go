package main

import (
	"fmt"
	"strconv"
	"strings"
)

// string 常用操作


func main06() {
	var str = "duanlei"
	contains := strings.Contains(str, "p") // 判断"duanlei" 中是否包含 "p", 包含返回 true 不包含返回 false
	fmt.Println(contains)

	var str1 = "bac,mkie,go"
	split := strings.Split(str1, ",") // 切割字符串
	fmt.Println(split)

	var str2 = "   bbbaabb    "
	space := strings.TrimSpace(str2)
	fmt.Println(space)

	fold := strings.EqualFold("A", "a") // 忽略大小写
	fmt.Println(fold)

	containsRune := strings.ContainsRune("duanlei", 'd') // 判断是否包含某个字符
	fmt.Println(containsRune)

	count := strings.Count("ddaabbddccnndd", "dd")
	fmt.Println(count)

	slices := []string{"abc", "hello", "mike", "go"}
	join := strings.Join(slices, "") //已什么方式组装字符串
	fmt.Println(join)

	// 字符串转换
	// Append : 将整数等 转换为字符串后，添加到现有的字节数组中
	//format ： 把其他类型转换为字符串
	//Parse ： 把字符串转化为其他类型

	slicess := make([]byte, 0, 1024)
	slicess = strconv.AppendBool(slicess, true)
	slicess = strconv.AppendQuote(slicess, "ajksfjskf")

	fmt.Println(string(slicess))

	var strs string
	strs = strconv.FormatBool(true)
	fmt.Println(strs)

	b, e := strconv.ParseBool("true")
	if e != nil {
		fmt.Println(e)
	}else {
		fmt.Println(b)
	}

}
