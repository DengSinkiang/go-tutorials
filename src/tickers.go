package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		//range 来迭代值每隔 500ms 发送一次的值
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	//2100ms停止打点器
	time.Sleep(time.Millisecond * 2100)
    	ticker.Stop()
    	fmt.Println("Ticker stopped")
}
