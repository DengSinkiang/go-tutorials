package main

import(
	"time"
	"fmt"
)

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)

    //<-Time.After 等待超时时间1秒后发送的值
    //如果这个操作超过了允许的1秒的话，将会执行超时 case
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}