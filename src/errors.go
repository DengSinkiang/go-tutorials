package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 45 {
		return -1, errors.New("can't work with 45")
	}
	return arg + 3, nil
}

type argError struct {
    arg  int
    prob string
}


//实现 Error 方法来自定义 error 类型
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 45 {
    	//使用 &argError 语法来建立一个 新的结构体
    	return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {
	for _, i := range []int{71, 45} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{72, 45} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }
    _, e := f2(45)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}