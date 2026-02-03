package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Amount     float64 `json:"amount"`
	CurrencyIN string  `json:"currency"`
	CurrencyTO string  `json:"to"`
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
	if req.CurrencyIN == "USD" && req.CurrencyTO == "INR" {
		answer = req.Amount * 90
	} else {
		fmt.Println("unsupported currency")
	}
	res := Result{Value: answer}
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func main() {
	http.HandleFunc("/convert", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
