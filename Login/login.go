package login

import (
	"encoding/json"
	"fmt"
	"golib/postgresql"
	"net/http"
)

//giriş yap

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
	rows, _ := db.Query("SELECT * FROM Userr")
	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, user.Email, user.Password)
		users = append(users, user)

	}
	//for _, item := range users {
	if _email == user.Email && _password == user.Password {
		fmt.Fprintf(w, "Login successful\n")
		fmt.Fprintln(w, "Hello", _firstName+" "+_lastname)
		peopleByte, _ := json.MarshalIndent(users, "", "\t")
		w.Write(peopleByte)
		//json.NewEncoder(w).Encode(item)
		return
	}
	//}
	//json.NewEncoder(w).Encode(&User{})
}
