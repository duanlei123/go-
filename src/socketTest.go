package main

import (
	"fmt"
	"net"
	"strings"
)

//socket

func main34() {
	// 监听
	listener, e := net.Listen("tcp", ":8000")
	if e != nil{
		fmt.Println("e = ", e)
		return
	}
	defer listener.Close()

	for{
		// 多客户端
		//阻塞等待用户连接
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("err = ", err)
			return
		}
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn){
	defer conn.Close()

	// 获取客户端地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("connect sucessful : ", addr)

	buf := make([]byte, 2048)

	for{
		//读取用户数据
		n, err := conn.Read(buf)
		if err !=nil{
			fmt.Println("e = ", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		if "exit" == string(buf[:n-2]){
			fmt.Println(addr," exit!")
			return
		}

		//将数据转换为大写，再发送回用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}