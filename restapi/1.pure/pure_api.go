package main

import (
	"fmt"
	"net/http"
)

func main () {
	fmt.Printf("Running server on Port 8080")
	
	mux := http.NewServeMux()

	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Main")
	})

	mux.HandleFunc("/badass", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Darth Vader")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error)
	}

}