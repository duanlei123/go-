package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main36() {
	// 监听
	listener, e := net.Listen("tcp", ":8000")
	if e != nil{
		fmt.Println("net.Listen err = ", e)
		return
	}

	defer listener.Close()

	//阻塞等待用户连接
	conn, e := listener.Accept()
	if e != nil{
		fmt.Println("listener.Accept err = ", e)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	n, e = conn.Read(buf) // 读取对方发送的文件名
	if e != nil{
		fmt.Println("conn.Read = ", e)
		return
	}
	fileName := string(buf[:n])

	//回复ok
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(fileName, conn)
}

func RecvFile(fileName string, conn net.Conn) {
	// 新建文件

	file, e := os.Create(fileName)
	if e != nil{
		fmt.Println("os.Create = ", e)
		return
	}

	buf := make([]byte, 1024*4)
	// 接收多少就写多少
	for{
		n, e := conn.Read(buf)  // 接收文件内容
		if e != nil{
			if e == io.EOF{
				fmt.Println("文件接收完毕")
			}else {
				fmt.Println("conn.Read err = ", e)
			}
			return
		}
		if n == 0{
			fmt.Println("文件接收完毕")
			return
		}
		file.Write(buf[:n]) //往文件写入内容
	}
}


