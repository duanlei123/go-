package calc

// 不同目录，包名不同
// 调用不同包里的函数： 格式： 包名.函数名
// 函数名大写Pulic.
//func init() {
//	fmt.Println("calc包中 init()函数执行")
//}func init() {
//	fmt.Println("calc包中 init()函数执行")
//}

func Add(a,b int) int{
	return a + b
}

func Minus(a,b int) int{
	return a-b
}