package main

import "fmt"

func main() {

	var a [6]int
    fmt.Println("array:", a)

    //指定下标位置的值
    a[3] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[3])

    fmt.Println("len:", len(a))

    b := [4]int{1, 2, 3, 4}
    fmt.Println("b:", b)

    var arrays [3][2]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 2; j++ {
            arrays[i][j] = i + j
        }
    }
    fmt.Println("arrays: ", arrays)


}