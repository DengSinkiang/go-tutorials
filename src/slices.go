package main

import "fmt"

func main() {

    //定义一个长度为4的 string 类型切片
    s := make([]string, 4)
    fmt.Println("s:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    //len 返回 slice 的长度
    fmt.Println("len:", len(s))

    //append 追加一个或多个新值的 slice，并返回
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("append:", s)

    //将 s 复制给 c
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    //得到 s[2],s[3],s[4]
    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("t:", t)

    //注意区别多维数组，因为这里是多维数据结构
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)


}
