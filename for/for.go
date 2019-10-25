package main

import (
	"fmt"
)

func main() {
	i := 1

	fmt.Println("For loop that prints 1-4:")
	for i <= 4 {
		iter := fmt.Sprint(i, " ")
		fmt.Print(iter)
		i = i + 1
	}

	fmt.Println("\n\nFor loop that prints 7-10:")
	for j := 7; j <= 10; j++ {
		iter := fmt.Sprint(j, " ")
		fmt.Print(iter)
	}

	fmt.Println("\n\nA potential infinite for-loop:")
	for {
		fmt.Println("loop\n")
		break
	}

	fmt.Println("For loop prints odd numbers 1-20:")
	for n := 0; n <= 20; n++ {
		if n%2 == 0 {
			continue
		}

		iter := fmt.Sprint(n, " ")
		fmt.Print(iter)
	}
}
