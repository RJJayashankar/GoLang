package main

import (
	"fmt"
)

type Result struct {
	Value float64 `json:"result"`
}

func main() {
	var size int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&size)
	var a = make([]int, size)

	for j := 0; j < size; j++ {
		fmt.Println("enter the elements of the array")
		fmt.Scanln(&a[j])

	}
	fmt.Println(a)
	var sum int

	for i := 0; i < size; i++ {
		sum += a[i]
	}
	var average int
	average = sum / size
	fmt.Println("the average of the array is:", average)
}
