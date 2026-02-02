package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("empty array", a)

	//setter
	a[4] = 100
	fmt.Println(a)

	//getter
	fmt.Println("get array index 4:", a[4])

	//length
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array b:", b)

}
