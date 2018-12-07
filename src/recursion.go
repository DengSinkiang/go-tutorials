package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	//递归调用函数
	return n * fact(n - 1)
}

func main() {
	fmt.Println(fact(5))
}