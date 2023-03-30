package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World! This is my first Go web app!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
