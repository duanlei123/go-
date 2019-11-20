package main

import (
	"fmt"
	"net"
)

// http 服务器开发
func main39() {
	// 监听
	listener, e := net.Listen("tcp", ":8000")
	if e != nil{
		fmt.Println("net.Listen err = ", e)
		return
	}
	defer  listener.Close()

	// 阻塞等待客户端连接
	conn, e := listener.Accept()
	if e != nil{
		fmt.Println("istener.Accept err = ", e)
		return
	}
	defer conn.Close()

	//读取用户信息
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	if n == 0{
		fmt.Println("conn.Read err = ", e)
		return
	}
	fmt.Println(string(buf[:n]))

}
