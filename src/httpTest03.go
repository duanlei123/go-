package main

import (
	"fmt"
	"net/http"
)

func main41() {

	// 发送http请求
	resp, err := http.Get("http://www.baidu.com")
	if err != nil{
		fmt.Println("http.GET err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	//fmt.Println(resp.Body)

	buf := make([]byte, 1024)
	var tmp string

	for{
		n, err := resp.Body.Read(buf)
		if n == 0{
			fmt.Println("err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println(tmp)

}
