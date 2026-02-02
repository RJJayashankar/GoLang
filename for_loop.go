package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
	}

	i := 1
	for i <= 10 {
		fmt.Println("index starting from", i)
		i = i + 1
	}

	for i := range 3 {
		fmt.Println("range", i)

	}

	for i := range 6 {
		if i%2 == 0 {
			continue

		}
		fmt.Println(i)
	}
}
