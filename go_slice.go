package main

import (
	"fmt"
)

func main() {
	var a []string
	if a == nil {
		fmt.Println("empty")
	}

	a = make([]string, 3)
	fmt.Println("length:", len(a), "capacity:", cap(a))

	a[0] = "a"
	a[1] = "b"
	a[2] = "c"

	fmt.Println(a)

	b := a[1:]
	fmt.Println(b)

	c := a[:2]
	fmt.Println(c)

}
