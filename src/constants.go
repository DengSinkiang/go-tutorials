package main

import (
	"fmt"
	"math"
)
//const 声明一个常量
const c string = "const"

func main() {
	fmt.Println(c)

	const a = 200
	const b = 2e10 / a
	fmt.Println(b)
	//类型转化
	fmt.Println(int64(b))
	//math.Cos 需要一个float64参数
	fmt.Println(math.Cos(a))
}