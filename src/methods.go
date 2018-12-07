package main

import "fmt"

type rect struct {
    width, height int
}

//指针类型接收
func (r *rect) area() int {
    return r.width * r.height
}

//值类型接收
func (r rect) perim() int {
    return 2 * (r.width + r.height)
}

func main() {
	
    r := rect{width: 102, height: 53}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}