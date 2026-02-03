package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type addition struct {
	num1 int
	num2 int
}

func (a addition) sum() int {
	return a.num1 + a.num2
}

func main() {
	fmt.Println(Person{name: "Jayashankar", age: 21})
	fmt.Println(addition{num1: 1, num2: 2}.sum())
}
