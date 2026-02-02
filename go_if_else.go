package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20
	var c int = 30
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)

	if a < b && a < c {
		fmt.Println("a is the smallest")
	} else if b < a && b < c {
		fmt.Println("b is the smallest")
	} else {
		fmt.Println("c is the smallest")
	}
}
