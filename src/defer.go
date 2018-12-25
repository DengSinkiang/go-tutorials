package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {

	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {

	fmt.Println("writing")
	fmt.Println(f, "data")
}

func closeFile(f *os.File) {

	fmt.Println("closing")
	f.Close()
}

func main() {

	f := createFile("./defer.txt")
	//Defer 被用来确保一个函数调用在程序执行结束前执行
	defer closeFile(f)
	writeFile(f)
}