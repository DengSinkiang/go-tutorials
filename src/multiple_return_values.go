package main

import "fmt"

//多个返回值
func vals() (int, int) {
	return 33, 73
}

func main() {

	//接收不同的返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//接收部分返回值
	_, c := vals()
	fmt.Println(c)
}
