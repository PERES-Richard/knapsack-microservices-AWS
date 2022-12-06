package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"knapset-generator/generator"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/generate", generateNewBag).Queries("bagSize", "{bagSize:[0-9,]+}", "numberOfItems", "{numberOfItems:[0-9,]+}").Methods("POST")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "generator is healthy\n")
	}).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: c.Handler(r),
	}

	println("Starting listening on port 8080...")
	log.Fatal(srv.ListenAndServe())
}

func generateNewBag(w http.ResponseWriter, r *http.Request) {
	bagSize, err := strconv.Atoi(r.FormValue("bagSize"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numberOfItems, err2 := strconv.Atoi(r.FormValue("numberOfItems"))
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("New generate request for a bagSize of : %d and %d items.\n", bagSize, numberOfItems)
	res, err := generator.GenerateNewKnapSet(bagSize, numberOfItems)

	if err != nil {
		http.Error(w, "Cannot generate a knapset with args bagSize "+strconv.Itoa(bagSize)+" and numberOfItems "+strconv.Itoa(numberOfItems)+".\n", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}
