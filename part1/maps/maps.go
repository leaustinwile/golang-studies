package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 12

	fmt.Println("Map:    ", m)

	v1, prs2 := m["k1"]
	fmt.Println("v1:     ", v1)
	fmt.Println("Length: ", len(m))

	delete(m, "k2")
	fmt.Println("Map:    ", m)

	noexist, prs1 := m["k2"]
	fmt.Println("Prs1:   ", prs1)
	fmt.Println("Prs2:   ", prs2)
	fmt.Println("Noexist:", noexist)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map:    ", n)
}
