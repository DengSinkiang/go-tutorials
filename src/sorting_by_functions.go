package main

import "sort"
import "fmt"

type ByLength []string

//实现了 sort.Interface 的 Len，Less 和 Swap 方法
//就可以使用 sort 包的通用 Sort 方法
func (s ByLength) Len() int {
        return len(s)
}

func (s ByLength) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
        return len(s[i]) < len(s[j])
}

func main() {
        fruits := []string{"peachewew", "rrrana", "kiwiewe"}
        sort.Sort(ByLength(fruits))
        fmt.Println(fruits)
}
