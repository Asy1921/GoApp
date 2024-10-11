package main

import (
	"backend/db/sqlc"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)
	const dbDriver= "postgres"
	const dbSource="postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable"
func sayHi(w http.ResponseWriter, r *http.Request) {

	var store *sqlc.Queries
	ctx := context.Background()
	conn,err :=pgx.Connect(ctx, dbSource)
	store= sqlc.New(conn)
	account,err:=store.GetAccount(ctx,3)
	fmt.Println(err.Error())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
func main() {
	fmt.Println("Running on localhost:8080");
	handleRequests()
}
func handleRequests() {
	http.Handle("/helloGoApi", http.HandlerFunc(sayHi))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
