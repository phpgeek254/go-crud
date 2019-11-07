package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting all th users")
}
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating New User")
}
func DeleterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting the user")
}
func EditUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Editting the user")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating the user.")
}
