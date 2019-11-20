package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main37() {
	// 提示输入文件
	fmt.Println("请输入需要传输的文件:")
	var path string
	fmt.Scan(&path)

	//获取文件名
	info, e := os.Stat(path)
	if e != nil{
		fmt.Println("os.Stat err = ", e)
		return
	}
	info.Name() // 文件名
	//主动连接服务器
	conn, e := net.Dial("tcp", ":8000")
	if e != nil{
		fmt.Println("net.Dial err = ", e)
		return
	}

	defer  conn.Close()
	//给接收方发送文件名
	_, e = conn.Write([]byte(info.Name()))
	if e != nil{
		fmt.Println("conn.Write err = ", e)
		return
	}

	// 接收对方回复 如果回复ok, 说明对方准备好了，可以发文件
	var n int
	buf := make([]byte, 1024)
	n, e = conn.Read(buf)
	if e != nil{
		fmt.Println("conn.Read err = ", e)
		return
	}
	if "ok" == string(buf[:n]){
		//发送文件内容
		SendFile(path, conn)
	}
}

// 发送文件内容
func SendFile(path string, conn net.Conn){
	// 已自读方式打开文件
	file, e := os.Open(path)
	if e != nil{
		fmt.Println("os.Open err = ", e)
		return
	}
	defer file.Close()
	// 读文件内容， 读多少给接收方发多少
	buf := make([]byte, 1027*4)
	for{
		n, e := file.Read(buf)
		if e != nil{
			if e == io.EOF{
				fmt.Println("文件发送完毕")
			}else {
				fmt.Println("file.Read err = ", e)
			}
			return
		}
		//发送内容
		conn.Write(buf[:n])
	}
}