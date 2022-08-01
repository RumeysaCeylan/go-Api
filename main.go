package main

import (
	"fmt"
	login "golib/Login"
	signup "golib/Signup"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all users
	router.HandleFunc("/login", login.Login).Methods("GET")

	// Create a user
	router.HandleFunc("/signup", signup.Signup).Methods("POST")
	fmt.Println("Server at 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
