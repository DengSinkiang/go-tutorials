package main

import "fmt"

func f(from string) {

	for i := 1; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")
	//两个 Go 协程在独立的 Go 协程中异步运行
	go f("goroutine")

	go func(msg string) {

		fmt.Println(msg)

	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
