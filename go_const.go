package main

import (
	"fmt"
)

func main() {
	const c float64 = 32.0
	fmt.Println(c)

	const f float64 = 68.0
	fmt.Println(f)

	celcius := (f - c) * 5.0 / 9.0
	fmt.Println("f", f, "to c is =", celcius)
}
