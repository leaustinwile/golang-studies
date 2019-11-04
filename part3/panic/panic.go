package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("D:\\file")
	if err != nil {
		panic(err)
	}
}
