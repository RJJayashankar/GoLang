package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Request struct{
	Name string `json:"name"`
	Lang string `json:"lang"`
} 

type Result struct{
	Value string `json:"result"`
}

func handler(w http.ResponseWriter, r *http.Request){
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var answer string
	if req.Lang == "fr"{
		answer = "bonjour"+ req.Name
	} else if req.Lang == "es"{
		answer = "hola" + req.Name
	}else{
		answer = "hello" + req.Name

	}
	res := Result{Value: answer}
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(res)
}

func main(){
	http.HandleFunc("/greet", handler)
	fmt.Println("Strting on port 8080")
	http.ListenAndServe(":8080", nil)
}