package signup

import (
	"encoding/json"
	"fmt"
	valid "golib/IsValid"
	"golib/postgresql"

	"net/http"
	"regexp"
)

//kayıt olmak
var firstname, lastname, email, password string

type User struct {
	Id        int
	firstName string
	lastName  string
	Email     string
	Password  string
}
type Login struct {
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

func Signup(w http.ResponseWriter, r *http.Request) {
	OpenConnention := postgresql.OpenConnention()
	var response = JsonResponse{}
	db := OpenConnention
	r.ParseForm()
	var login Login
	login.Password = r.FormValue("password")
	login.Email = r.FormValue("email")
	rows, _ := db.Query("SELECT * FROM User")
	for rows.Next() {
		rows.Scan(&user.Id, &user.firstName, &user.lastName, &user.Email, &user.Password)
	}
	if login.Email == "" || password == "" {
		fmt.Fprintf(w, "cannot be empty")
	} else {
		if user.Email == login.Email {
			fmt.Fprintf(w, "email is used")
		} else {
			if CheckEmail(login.Email) == true {
				db.Exec("INSERT INTO User(firstName,lastName,email,password) VALUES($1,$2,$3,$4)", login.firstName, login.lastName, login.Email, login.Password)

				peopleByte, _ := json.MarshalIndent(login, "", "\t")

				w.Header().Set("Content-Type", "application/json")

				w.Write(peopleByte)

				defer db.Close()

				valid.IsValid(w, r)
				response = JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}

			}
		}
	}
	json.NewEncoder(w).Encode(response)
}
func CheckEmail(mail string) bool {
	match, _ := regexp.MatchString("[^@]+@[^@]+\\.[^@]+", mail)
	return match
}
