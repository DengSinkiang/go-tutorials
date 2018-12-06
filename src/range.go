package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range 在 map 中迭代键值对
	values := map[string]string{"a": "boy", "b": "girl"}
	for k, v := range values {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range 可以只遍历map的键
	for k := range values {
        fmt.Println("key:", k)
    }

    for i, c := range "ab" {
        fmt.Println(i, c)
    }

}