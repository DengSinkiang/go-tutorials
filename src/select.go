package main

import (
	"time"
	"fmt"
)

//管道选择器 让你可以同时等待多个管道操作

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "deng"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "xin"
	}()

	for i := 0; i < 2; i++ {
		//select 关键字来同时等待这两个值
        select {
        case msg1 := <-c1:
            	fmt.Println("received", msg1)
        case msg2 := <-c2:
            	fmt.Println("received", msg2)
        }
    }

}
