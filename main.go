package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homepage")
}
func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:5001", nil)
}
