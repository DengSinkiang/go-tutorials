
package main

import "fmt"

func main() {

	var a = "hello"
	fmt.Println(a)

	//var声明一个或多个变量
	var b, c int = 100, 200
	fmt.Printf("b=%d c=%d\n", b, c)

	var d = false
	fmt.Println(d)
	//声明后却没有给出对应的初始值时，变量将会初始化为默认值。
	var e string
	fmt.Printf("e:%s\n", e)
	
	//:= 语法是声明并初始化变量的简写
	f := 3.4
	fmt.Println(f)

}