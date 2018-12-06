package main

import "fmt"

//inSeq 返回在其内部定义的匿名函数 该返回的函数采用闭包的方式
func intSeq() func() int {
	i := 10
	return func() int {
		i++
		return i
	}
}

func main() {

	//调用 ineSeq 函数将返回值赋值给 nextInt 该函数的值包含本身的值 i
	//每次调用nextInt都会更新 i
	nextInt := intSeq()

	//多次调用 nextInt 查看闭包的值的变化
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//该状态对于特定函数是唯一的
	newInts := intSeq()
	fmt.Println(newInts())
}
