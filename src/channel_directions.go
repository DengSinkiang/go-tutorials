package main

import "fmt"

//chan<- 只允许发送数据	
func send(pings chan<- string, msg string) {
	pings <- msg
}


//<-chan 只允许接受数据
func receive(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <-msg
}

func main() {
    	pings := make(chan string, 1)
    	pongs := make(chan string, 1)
    	send(pings, "passed message")
    	receive(pings, pongs)
    	fmt.Println(<-pongs)
}
