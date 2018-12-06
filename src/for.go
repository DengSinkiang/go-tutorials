package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("********")
	for j := 5; j <= 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("********")
	for {
		fmt.Println("break")
		//break 终止循环
		break
	}
	
	fmt.Println("********")
	for k := 1; k <= 10; k++ {
		if k % 2 == 0 {
			//跳出本次循环
			continue
		}
		fmt.Println(k)
	}

}