package main

import "fmt"

func sum(nums ...int) {

	//多个 int 作为参数
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {

	//变参函数调用
	sum(13, 22)
	sum(11, 22, 33)

	//slice 调用 func(slice...)
	nums := []int{31, 32, 3, 4}
	sum(nums...)
}
