package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type Result struct {
	Sum int `json:"sum"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req Request
	
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sum := req.Num1 + req.Num2
	res := Result{Sum: sum}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func main() {
	http.HandleFunc("/sum", handler)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
