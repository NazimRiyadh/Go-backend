package main

import (
	"fmt"
	"net/http"
)

func about_handler(w http.ResponseWriter, r *http.Request) { // handler function for route
	fmt.Fprintf(w, "This is about page")
}

func profile_handler(w http.ResponseWriter, r *http.Request) { // handler function for route
	fmt.Fprintf(w, "This is profile page")
}

func main() {
	mux := http.NewServeMux() //router creation

	mux.HandleFunc("/about", about_handler) //route creation
	mux.HandleFunc("/profile", profile_handler)

	fmt.Println("Starting server at :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil { //checking if there is an error
		fmt.Println("Error starting server:", err)
	}

}
