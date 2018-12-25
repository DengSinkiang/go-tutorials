package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"f", "s", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{9, 3, 1, 6}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	//检查是否排好序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}