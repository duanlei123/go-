package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// go 语言文件操作

// 文件分类: 设备文件
// 				屏幕(标准输出设备) fmt.Println() 向标准输出设备写内容
// 				键盘(标准输入设备) fmt.Scan()     从标准输入设备读内容
// 			磁盘文件
//				放在存储设备上的文件
//				文本文件： 已记事本打开，能看到内容(不是乱码)
//				二进制文件： 用记事本打开，看到乱码(比如图片.....)

//为什么需要文件？
//	持久化数据

// 文件的读写
func main26() {
	path := "./demo.txt"
	//WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)
}

//每次读取一行
func ReadFileLine(path string) {
	// 打开文件
	file, e := os.Open(path)
	if e != nil{
		fmt.Println("e = ", e)
		return
	}
	//关闭文件
	defer file.Close()
	// 新建一个缓存区
	reader := bufio.NewReader(file)

	for{
		bytes, e := reader.ReadBytes('\n')
		if e != nil{
			if e == io.EOF{
				break
			}
			fmt.Println("err = ", e)
		}
		fmt.Println(string(bytes))
	}
}

func ReadFile(path string) {
	// 打开文件
	file, e := os.Open(path)
	if e != nil{
		fmt.Println("e = ", e)
		return
	}
	//关闭文件
	defer file.Close()

	buf := make([]byte, 1024*2)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF{
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(string(buf[:n]))
}

func WriteFile(path string) {
	// 打开文件/新建文件
	file, err := os.Create(path)
	if err != nil{
		fmt.Println("err = ", err)
		return
	}
	// 使用完关闭文件
	defer file.Close()

	var buf string
	for i := 0; i < 10; i++{
		// 将字符串存储在buf中
		buf = fmt.Sprintf("i = %d\n", i)
		writeString, err := file.WriteString(buf)
		if err != nil{
			fmt.Println("err = ", err)
		}
		fmt.Println("writeString = ", writeString)
	}
}



