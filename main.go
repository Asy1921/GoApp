package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello from Go Backed")
}
func main() {
	fmt.Println("Running on localhost:8080");
	handleRequests()
}
func handleRequests() {
	http.Handle("/helloGoApi", http.HandlerFunc(sayHi))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
