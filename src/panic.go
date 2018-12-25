package main

import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/data")
	if err != nil {
		panic(err)
	}
}