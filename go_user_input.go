package main

import (
	"fmt"
)

var name string
var age int

func main() {
	fmt.Println("Enter your name")
	fmt.Scanln(&name)

	fmt.Println("Enter your age")
	fmt.Scanln(&age)

	fmt.Println("Your name is", name)
	fmt.Println("Your age is", age)

}
