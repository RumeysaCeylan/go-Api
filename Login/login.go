package login

import (
	"encoding/json"
	"fmt"
	"golib/postgresql"
	"net/http"
)

//giriş yap

type User struct {
	Id        int
	firstName string
	lastName  string
	Email     string
	Password  string
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}

var user User

func Login(w http.ResponseWriter, r *http.Request) {
	openConnention := postgresql.OpenConnention()

	var users []User
	db := openConnention
	defer db.Close()
	r.ParseForm()
	_firstName := r.FormValue("firstName")
	_lastname := r.FormValue("lastName")
	_password := r.FormValue("password")
	_email := r.FormValue("email")
	rows, _ := db.Query("SELECT * FROM User")
	for rows.Next() {
		rows.Scan(&user.Id, &user.firstName, &user.lastName, user.Email, user.Password)
		users = append(users, user)

	}
	if _email == user.Email && _password == user.Password {
		fmt.Fprintf(w, "Login successful\n")
		fmt.Fprintln(w, "Hello", _firstName+" "+_lastname)
		peopleByte, _ := json.MarshalIndent(user, "", "\t")
		w.Write(peopleByte)
	}
	var response = JsonResponse{Type: "success", Data: users}

	json.NewEncoder(w).Encode(response)
}
