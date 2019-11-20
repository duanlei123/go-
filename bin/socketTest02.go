package main

import (
	"fmt"
	"net"
	"os"
)

func main35() {
	// 主动连接服务器
	conn, e := net.Dial("tcp", ":8000")
	if e != nil{
		fmt.Println("e = ", e)
		return
	}
	defer conn.Close()

	// 接收服务器回复信息
	go func() {
		buf := make([]byte, 1024)
		for{
			n, e := conn.Read(buf)
			if e != nil{
				fmt.Println("e = ", e)
				return
			}
			fmt.Println("回复：",string(buf[:n]))
		}
	}()
	//键盘输入
	for{
		str := make([]byte, 1024)
		n, e := os.Stdin.Read(str) // 接盘读取内容
		if e != nil{
			fmt.Println("e = ", e)
			return
		}
		// 发送数据
		conn.Write(str[:n])
	}
}
