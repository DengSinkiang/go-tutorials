package main

import (
	"time"
	"fmt"
)

func main() {

	//定时器等待2秒
	timer1 := time.NewTimer(time.Second * 2)

	//<-timer1.C 直到这个定时器的通道 C 明确的发送了定时器失效的值之前 将一直阻塞
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second * 1)
    	go func() {
        	<-timer2.C
        	fmt.Println("Timer 2 expired")
    	}()
    
    	stop2 := timer2.Stop()
    	if stop2 {

        	fmt.Println("Timer 2 stopped")
    	}
}
