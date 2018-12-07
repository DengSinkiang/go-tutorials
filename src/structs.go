package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {

	fmt.Println(person{"a", 20})

	fmt.Println(person{name: "b", age: 21})

	fmt.Println(person{name: "c"})

	fmt.Println(&person{"d", 22})

	s := person{name: "e", age: 23}
	fmt.Println(s.name)

	//将 s 的地址赋值给 sp，指向同一个地址
	sp := &s
	fmt.Println(sp.name)

	sp.name = "dxj"
	fmt.Println(sp.name)

}