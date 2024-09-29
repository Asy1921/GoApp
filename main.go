package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello from Go Backed")
}
func main() {
	handleRequests()
}
func handleRequests() {
	http.Handle("/helloGoApi", http.HandlerFunc(sayHi))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
