package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"knapset-generator/generator"
	"net/http"
)

type GeneratorRequest struct {
	BagSize       int `json:"bagSize"`
	NumberOfItems int `json:"numberOfItems"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/generate", generateNewBag).Methods("POST")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "healthy")
	}).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	println("Starting listening on port 8080...")
	srv.ListenAndServe()
}

func generateNewBag(w http.ResponseWriter, r *http.Request) {
	var req GeneratorRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := generator.GenerateNewKnapSet(req.BagSize, req.BagSize)

	json.NewEncoder(w).Encode(res)
}
