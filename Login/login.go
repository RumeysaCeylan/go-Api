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

	_password := r.FormValue("password")
	_email := r.FormValue("email")
	rows, _ := db.Query("SELECT * FROM Userr")
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r) // Gets params
	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, user.Email, user.Password)
		users = append(users, user)

	}
	for _, item := range users {
		if _email == item.Email || _password == item.Password {

			//if string(item.Id) == params["id"] {
			fmt.Fprintf(w, "Login successful\n")

			//peopleByte, _ := json.MarshalIndent(user, "", "\t")
			//w.Write(peopleByte)
			json.NewEncoder(w).Encode(users)
			return
			//}

		}
	}
	//json.NewEncoder(w).Encode(&User{})
}
