package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// 爬取百度贴吧 小案例
func main() {
	var start , end int
	fmt.Println("请输入起始页(>= 1):")
	fmt.Scan(&start)

	fmt.Println("请输入终止页(>= 起始页):")
	fmt.Scan(&end)

	DoWork(start, end)
}

func DoWork(start int, end int) {
	fmt.Printf("正在爬取 %d - %d 的页面\n", start, end)

	//明确目标
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100

	page := make(chan int)
	for i := start; i <= end; i++{
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++{
		fmt.Printf("第%d页面爬取完成\n", <-page)
	}
}

//
func SpiderPage(i int, page chan<- int)  {
	url := "http://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa((i-1) * 50)
	fmt.Printf("正在爬第%d网页 %s\n", i, url)
	//爬取
	result , err := httpGet(url)
	if err != nil{
		fmt.Println("httpGet err = ", err)
		return
	}

	// 把爬取的内容，写入一个文件
	fileName := strconv.Itoa((i))+".html"
	file, err := os.Create(fileName)
	if err != nil && err != io.EOF{
		fmt.Println("os.Create err = ", err)
		return
	}
	file.WriteString(result) // 写内容
	defer file.Close() // 关闭文件

	page <- i
}

//爬取网页内容
func httpGet(url string) (result string, err error) {
	resp,err1 := http.Get(url)
	if err1 != nil{
		err = err1
		return
	}

	buf := make([]byte, 1024*4)
	for{
		n, _ := resp.Body.Read(buf)
		if n == 0{ // 读取结束，或者出问题
			//fmt.Println("err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

