package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The API is working fine")
}
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexHandler).Methods("POST")

	// User routes definition.
	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/users/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/users/{name}/{email}", EditUser).Methods("PUT")
	router.HandleFunc("/users/{name}", DeleterUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1200", router))
}

func main() {
	fmt.Println("Starting the ORM TUTORIAL SERVER")
	InitialMigration()
	handleRequests()
}
