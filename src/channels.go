package main

import "fmt"

func main() {

	//make 创建管道
	messages := make(chan string)
	//channel<- 语法 发送(send)一个新的值到通道中
	//<-channel 语法 从通道中接收(receives)一个值
	go func() { messages <- "deng" }()

	msg := <-messages
    fmt.Println(msg)
}