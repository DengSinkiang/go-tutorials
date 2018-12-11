package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	go func() {
		//循环的从 jobs 接收数据
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				//接收完毕通过 done 管道进行通知
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	end := <-done
	fmt.Println(end)
}
