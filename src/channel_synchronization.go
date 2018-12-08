package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
func main() {

	done := make(chan bool, 1)
	//运行一个 worker Go协程，并给予用于通知的通道
	go worker(done)
	<-done
}
