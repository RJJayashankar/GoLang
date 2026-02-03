package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Num1      int    `json:"num1"`
	Num2      int    `json:"num2"`
	Operation string `json:"operation"`
}
type Result struct {
	Value int `json:"result"`
}


func handler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	var answer int

	switch req.Operation {
	case "+":
		
		answer = req.Num1 + req.Num2

	case "*":
		answer = req.Num1 * req.Num2

	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	res := Result{Value: answer}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/calculate", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}