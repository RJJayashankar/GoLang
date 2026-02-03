package main

import (
	"encoding/json"
	"fmt"
)

type numbers struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type result struct {
	Sum int `json:"sum"`
}

func main() {
	nums := numbers{Num1: 10, Num2: 20}
	total := nums.Num1 + nums.Num2
	res := result{Sum: total}
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
	}
	fmt.Println(string(data))
}
