package main

import "fmt"

func intSeq() func() int {
	i := 10
	return func() int {
		i++
		return i
	}
}

func main() {

	//调用 ineSeq 函数将返回值赋值给 nextInt
	nextInt := intSeq()

	//多次调用 nextInt 查看闭包的值的变化
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//该状态对于特定函数是唯一的
	newInts := intSeq()
	fmt.Println(newInts())
}
