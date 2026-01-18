		package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var products []product

type product struct {
	Id      int
	Name    string
	Price   float64
	Img_url string
}

func get_products_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Allow-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json") //for json response
	w.Header().Set("Access-Control-Allow-Origin", "*") //CORS policy
	if r.Method != "GET" {
		http.Error(w, "plz provide GET method", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}

func add_product_handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Allow-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*") //CORS policy
	w.Header().Set("Content-Type", "application/json") //for json response

	if r.Method != "POST" {
		http.Error(w, "plz provide POST method", 400)
		return
	}
	var new_product product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&new_product)
	if err != nil {
		http.Error(w, "Error decoding JSON data", 400)
		return
	}

	new_product.Id = len(products) + 1
	products = append(products, new_product)

	encoder := json.NewEncoder(w)
	encoder.Encode(new_product)

}

func main() {
	mux := http.NewServeMux() //router creation

	mux.HandleFunc("/products", get_products_handler) //route creation

	mux.HandleFunc("/add_product", add_product_handler) //route creation

	fmt.Println("Starting server at :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil { //checking if there is an error
		fmt.Println("Error starting server:", err)
	}

}

func init() {
	prod1 := product{1, "Laptop", 45000.00, "https://example.com/laptop.jpg"}
	prod2 := product{2, "Smartphone", 25000.00, "https://example.com/smartphone.jpg"}
	prod3 := product{3, "Headphones", 5000.00, "https://example.com/headphones.jpg"}
	products = append(products, prod1, prod2, prod3)
}
