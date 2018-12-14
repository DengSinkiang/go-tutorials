package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"runtime"
)

func main() {

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				//对计数器每隔 1ms 进行一次加一操作
				//AddUint64 来让计数器自动增加
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)

}