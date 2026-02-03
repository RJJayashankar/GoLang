package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Tell the server: "When someone visits the home page (/), run this function."
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		// 2. We send HTML (Big Text) to the browser
		fmt.Fprint(w, "<h1>Hello! This is my web page.</h1>")
	})

	// 3. Start the server on port 8080
	fmt.Println("Server running... go to http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
