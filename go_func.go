package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func multi(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(1, 1))
	fmt.Println(multi(1, 3))

}
