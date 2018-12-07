package main

import "fmt"

//值传递 zeroval 将从调用它的那个函数中得到一个 ival 形参的拷贝
func zeroval(ival int) {
	ival = 0
}

//引用传递 调用函数会改变该地址的值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 111
	fmt.Println("initval:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
}