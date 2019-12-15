package main

import "fmt"

func main() {
	fmt.Println("变量demo")
/*
	//var a int // var 开头，后置变量名 变量类型，结尾无须分号。
	//fmt.Println(a)//变量会初始附零值，此处输出0

	//var b int = 123 // 定义同时并赋值
	//fmt.Println(b)//此处输出123

 */
/*
声明多个变量
	var a,b int
	var (
		c int
		d string
		e []float32
		f func() bool
		g struct {
			x int
		}
	)
*/
/*
	//简化格式定义变量   名字 := 表达式
	//需要注意的是，简短模式（short variable declaration）有以下限制：
	//定义变量，同时显式初始化。
	//不能提供数据类型。
	//只能用在函数内部。
	a := 123
	b,c:=1, "abc"
	fmt.Println(a,b,c)

 */

	a, b := 10, 20

	//可以通过多重赋值进行交换
	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
