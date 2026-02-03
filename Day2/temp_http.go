package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Temp   float64 `json:"temp"`
	Symbol string  `json:"symbol"`
}

type Result struct {
	Value float64 `json:"result"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}
	var answer float64
	switch req.Symbol {
	case "C":
		answer = (req.Temp * 9 / 5) + 32.0
	case "F":
		answer = (req.Temp - 32.0) * 5 / 9

	default:
		fmt.Println("Invalid symbol")
	}
	res := Result{Value: answer}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/convert", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
