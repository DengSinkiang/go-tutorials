package main

import (

	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0

	for r := 0; r <= 1000; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				//Lock() 这个 mutex 来确保对 state 的独占访问 读取选定的键的值
				mutex.Lock()
				total += state[key]
				//Unlock() 这个 mutex，并且 ops 值加 1
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                //确保 Go 协程不会在调度中死亡 使用 runtime.Gosched() 进行释放
                runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()

}