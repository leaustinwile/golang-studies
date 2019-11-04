package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{7, 2, 4}
	strs := []string{"c", "a", "b"}

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s = sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
