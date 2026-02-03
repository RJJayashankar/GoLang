package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["india"] = "Delhi"
	m["USA"] = "washington DC"

	fmt.Println(m)

	n := map[string]int{
		"apple":      1,
		"samsung":    2,
		"blackberry": 3,
	}

	fmt.Println(n)

	n2 := map[string]int{
		"apple":      1,
		"samsung":    2,
		"blackberry": 3,
	}
	if maps.Equal(n, n2) {
		fmt.Println("map", n, "and", n2, "are equal")
	}

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a == b)
}
