package main

import "fmt"

func plus(a int, b int) int {

    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    res := plus(12, 23)
    fmt.Println("12+23 =", res)
    res = plusPlus(11, 22, 33)
    fmt.Println("11+22+33 =", res)
}