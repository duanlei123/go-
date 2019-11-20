package main

import (
	"fmt"
	"io"
	"os"
)

// 文件拷贝
func main27() {
	// 获取命令行参数
	list := os.Args
	if len(list) != 3{
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	// 只读方式打开原文件
	sF, e1 := os.Open(srcFileName)
	if e1 != nil{
		fmt.Sprintln("e = ", e1)
		return
	}

	// 新建目的文件d
	dF, e2 := os.Create(dstFileName)
	if e2 != nil{
		fmt.Sprintln("e = ", e2)
		return
	}

	// 关闭文件
	defer sF.Close()
	defer dF.Close()

	// 核心处理
	buf := make([]byte, 4*1024)
	for{
		n, err1 := sF.Read(buf) // 读
		if err1 != nil {
			if err1 == io.EOF{ //文件读取完毕
				break
			}
			fmt.Sprintln("err1 = ", err1)
		}
		// 写
		dF.Write(buf[:n])
	}
}
