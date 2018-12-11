package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	//不关闭会失锁 循环继续阻塞执行 等待接收第三个值
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
