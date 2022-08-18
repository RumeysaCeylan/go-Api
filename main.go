package main

import (
	"fmt"
	login "golib/Login"
	signup "golib/Signup"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	router := http.NewServeMux()

	// Route handles & endpoints

	// Get all users
	router.HandleFunc("/login", login.Login)

	// Create a user
	router.HandleFunc("/signup", signup.Signup)
	fmt.Println("Server at 8000")

	http.ListenAndServe(":8000", router)
}
