package main

import (
	"fmt"
	"time"
)

func main() {

	i := 3
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool type")
		case int:
			fmt.Println("I'm a int type")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(false)
	whatAmI(10)
	whatAmI("hello")
}
