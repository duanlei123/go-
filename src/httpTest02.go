package main

import (
	"fmt"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request)  {

	fmt.Println(req.Host)
	fmt.Println(req.Method)
	fmt.Println(req.Body)
	fmt.Println(req.Header)

	w.Write([]byte("hello world"))
}

func main40() {
	// 注册处理函数
	http.HandleFunc("/", HelloServer)
	//监听绑定
	http.ListenAndServe(":8000", nil)
}
