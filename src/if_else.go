package main

import "fmt"

func main() {

	if 5 % 2 == 0 {
		fmt.Println("5 is even")
	} else {
		fmt.Println("5 is odd")
	}

	if 4 % 2 == 0 {
        fmt.Println("4 is divisible by 2")
    }
    //在条件语句之前可以有一个声明语句,此变量可以在所有的条件分支中使用
    if number := 5; number < 0 {
        fmt.Println(number, "is negative")
    } else if number < 10 {
        fmt.Println(number, "has 1 digit")
    } else {
        fmt.Println(number, "has multiple digits")
    }
}