package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	fmt.Println("map:", m)

	a1 := m["a"]
	fmt.Println("a1:", a1)

	//len 返回键值对的数目
	fmt.Println("len:", len(m))

        //delete 移除相应的键值对
	delete(m, "a")
	fmt.Println("map:", m)

	//_表示忽略该值，ok 表示此键是否存在 map 中
	_, ok := m["a"] 
	if !ok {
		fmt.Println("不存在该值")
	}
	
	//fmt.Println 打印时，是以 map[k1:v1 k2:v2] 的格式输出的
	n := map[string]int{"foo": 1, "bar": 2}
        fmt.Println("map:", n)
}
