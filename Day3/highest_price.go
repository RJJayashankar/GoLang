package main

import (
	"fmt"
)

type crops struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type highest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	crops_data := []crops{
		{Name: "snake gourd", Type: "Vegetable", Price: 50.00},
		{Name: "bitter gourd", Type: "Vegetable", Price: 65.00},
		{Name: "Potato", Type: "Vegetable", Price: 25.00},
		{Name: "Brinjal", Type: "vegetable", Price: 32.00},
	}

	var highest_price highest
	highest_price.Name = ""
	highest_price.Price = 0.00

	for i := 0; i < len(crops_data); i++ {
		if crops_data[i].Price > highest_price.Price {
			highest_price.Price = crops_data[i].Price
			highest_price.Name = crops_data[i].Name
		}

	}
	fmt.Println(highest_price.Name, highest_price.Price)
}
