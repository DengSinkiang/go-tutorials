package main

import (
	"fmt"
	"math"
)

//几何体接口
type geometry interface {
    area() float64
    perim() float64
}

//长方形
type rect struct {
    width, height float64
}
//圆
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 6, height: 8}
    c := circle{radius: 10}

    measure(r)
    measure(c)
}