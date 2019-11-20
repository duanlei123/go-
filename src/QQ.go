package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//用户结构体
type Client struct {
	C chan string // 用户发送数据的管道
	Name string // 用户名
	Addr string // 地址
}

var onlineMap map[string]Client

var message = make(chan string)

func Manager() {
	onlineMap = make(map[string]Client)
	for{
		msg := <- message // 没有消息会阻塞
		//遍历map
		for _, cli := range onlineMap{
			cli.C <- msg
		}
	}
}

func WriteMagToClient(client Client, conn net.Conn) {
	for msg := range  client.C{
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name +": " + msg
	return buf
}

func HandleConn1(conn net.Conn)  {
	defer conn.Close()
	// 获取客户端信息
	cliAddr := conn.RemoteAddr().String()
	//创建结构体
	cli := Client{make(chan string), cliAddr,cliAddr}
	// 把结构体加入到map
	onlineMap[cliAddr] = cli
	//新开一个协程，专门给当前客户端发送信息
	go WriteMagToClient(cli, conn)
	//广播某个人在线
	message <- MakeMsg(cli, "login")
	//提示我是谁
	cli.C <- MakeMsg(cli, "I am here")

	isQuit := make(chan bool) // 客户端是否退出

	hasData := make(chan bool) // 对方是否有数据发送

	// 新开一个go 接收用户发送的数据
	go func() {
		buf := make([]byte, 2048)
		for{
			n, err := conn.Read(buf)
			if n == 0{ // 对方断开，或者出问题
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}
			// 转发此内容
			msg := string(buf[:n-1])

			if len(msg) == 3 && msg == "who"{
				//遍历map, 给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap{
					msg = tmp.Addr + ":"+ tmp.Name +"\n"
					conn.Write([]byte(msg))
				}
			}else if len(msg) >= 8 && msg[:6] == "rename"{
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok:\n"))
			}else {
				message <- MakeMsg(cli, msg)
			}
			hasData <- true
		}
	}()
	for{
		//通过select 检测channel 的流动
		select {
		case <-isQuit:
			delete(onlineMap,cliAddr) // 当前用户从map移除
			message <- MakeMsg(cli,"login out!") // 广播下线
			return
		case <-hasData:

		case <-time.After(60 * time.Second):
			delete(onlineMap,cliAddr) // 当前用户从map移除
			message <- MakeMsg(cli,"time out leave out !") // 广播下线
			return
		}
	}
}

func main38() {
	// 监听
	listener, e := net.Listen("tcp", ":8000")
	if e != nil{
		fmt.Println("net.Listen err = ", e)
		return
	}
	defer  listener.Close()
	// 新开一个协程，转发消息， 只要有消息来了，遍历map, 给每个成员发送消息
	go Manager()
	// 主协程, 循环阻塞等待用户连接
	for{
		conn, e := listener.Accept()
		if e != nil{
			fmt.Println("listener.Accept err = ", e)
			continue
		}
		go HandleConn1(conn) // 处理用户连接
	}
}
