package main

// 正则表达式

// . 匹配任意任意除 “\n” 换行字符 在DOTALL模式中也能匹配换行符   a.b  -  abc
// \ 转义字符, 使用一个字符改变原来的意思，如果字符串中有字符“*”需要匹配，可以使用\*或者字符集[*]  a\.b 或 a\\b  - a.b a\b
// [...] 字符集(字符类),对应的位置可以说字符集中任意的字符,字符集中的字符可以逐一列出，也可以给出范围 如[abc] 或 [a-c] 第一次字符如果是 "^" 表示取反
//        如：[^abc]:表示不是abc的其他字符, 所有的特殊字符在字符集中都失去了特殊含义，字符集中如果使用 ]、- 或 ^ 可以在前面加一个反斜杠,或者把 ] 、-
//			  放在第一个字符，把 ^ 放在非第一字符 a[bcd]e - abe ace ade

// \d 数字 [0-9] a\dc - a1c a2c a3c ...
// \D 非数字,字符集[^\d] a\Dc - abc csc ...
// \s 空白字符: [<空格>\t\r\n\f\v]     a\sc - a c
// \S 非空白字符 [^\s]  a\Sc - abc...
// \w 单词字符 [A-Za-z0-9] a\wb - abc
// \W 非单词字符 [^\w] a\Wb- a c
// * 匹配前一个字符0 或 无线次  ab* - ab 或 abbbbbb
// + 匹配前一个字符 1 或无线次  ab+ - abb 或 abbbbb
// ? 匹配前一个字符 0 次或 1次  ab? = a 或 ab
// {m} 匹配前一个字符 m 次  ab[2]c - abbc
// {m,n} 匹配前一个字符 m-n 次 m 和 n 可以省略 省略m 0-n 省略 n m-无限次  ab[1,2]c - abc abbc
//

import (
	"fmt"
	"regexp"
)

func main24() {
	buf := "abc  azc a7c aac 888 a9c tac"
	// 1, 解释规则
	compile := regexp.MustCompile(`a\dc`)
	if compile == nil{
		fmt.Println("err = ", compile)
		return
	}
	// 2, 根据规则提取信息
	submatch := compile.FindAllStringSubmatch(buf, -1)
	fmt.Println(submatch)

	buf1 := "3.14 567 agsdg 1.23 7. 8 8.99 lsdlal 6.66 33.56"
	// 使用正则表达式匹配buf1中的有效小数
	mustCompile := regexp.MustCompile(`\d+\.\d+`)
	if mustCompile == nil {
		fmt.Println("err = ", mustCompile)
		return
	}
	stringSubmatch := mustCompile.FindAllStringSubmatch(buf1, -1)
	fmt.Println(stringSubmatch)


	buf2 := `
		<!DOCTYPE html>
		<html lang="zh-CN">
		<head>
			<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
			<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
			<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
			<meta charset="utf-8">
			<link rel="shortcut icon" href="/static/img/go.ico">
			<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
			<meta name="author" content="polaris <polaris@studygolang.com>">
			<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
			<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
		</head>
			<div>数据</div>
			水电费
			洛凯股份
			<div>偶尔会</div>
			<div>大姐夫</div>
		<frameset cols="15,85">
			<frame src="/static/pkgdoc/i.html">
			<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
			<noframes>
			</noframes>
		</frameset>
		</html>
	`
	mustCompile1 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if mustCompile1 == nil {
		fmt.Println("err = ", mustCompile1)
		return
	}
	stringSubmatch1 := mustCompile1.FindAllStringSubmatch(buf2, -1)
	fmt.Println(stringSubmatch1)

	// 过滤 <></>
	for _, text := range stringSubmatch1{
		fmt.Println(text[1])
	}
}