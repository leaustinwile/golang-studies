package main

import (
	"fmt"
)

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0", 7.0/3.0)

	fmt.Println(true && false)   // AND
	fmt.Println(true || true)    // OR
	fmt.Println(!true)           // NOT
	fmt.Println(!true || !false) // XOR
}
