package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID       string
	Name     string
	Quantity int
	Price    float64
}

var products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage!")
	log.Println("Endpoint Hit: homepage")
}
func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(products)
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: getProduct")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, product := range products {
		if string(product.ID) == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	http.ListenAndServe("localhost:5001", myRouter)
}
func main() {
	products = []Product{
		{ID: "1", Name: "Chair", Quantity: 10, Price: 19.99},
		{ID: "2", Name: "Table", Quantity: 5, Price: 29.99},
		{ID: "3", Name: "Sofa", Quantity: 20, Price: 9.99},
	}
	handleRequests()
}
