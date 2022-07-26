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
//var firstname, lastname, email, password string

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Login struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

var user User
var lg []Login

func Signup(w http.ResponseWriter, r *http.Request) {

	OpenConnention := postgresql.OpenConnention()

	db := OpenConnention
	r.ParseForm()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM Userr")
	var login Login
	login.FirstName = r.FormValue("firstName")
	login.LastName = r.FormValue("lastName")
	login.Email = r.FormValue("email")
	login.Password = r.FormValue("password")

	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	}
	if login.Email == "" || login.Password == "" {
		fmt.Println("cannot be empty")
		fmt.Fprintf(w, "cannot be empty")
	} else {
		if user.Email == login.Email {
			fmt.Fprintf(w, "Email is used")
			fmt.Println("email is used")

		} else {
			if CheckEmail(login.Email) && valid.IsValid(w, r) {
				db.Exec("INSERT INTO Userr(firstname,lastname,email,password) VALUES($1,$2,$3,$4)", login.FirstName, login.LastName, login.Email, login.Password)
				peopleByte, _ := json.MarshalIndent(user, "", "\t")

				w.Header().Set("Content-Type", "application/json")

				w.Write(peopleByte)

				defer db.Close()
				//_ = json.NewDecoder(r.Body).Decode(&login)

				lg = append(lg, login)
				//json.NewEncoder(w).Encode(login)

			} else {
				fmt.Fprintln(w, "record failed error insert!! ")
			}
		}

	}

}
func CheckEmail(mail string) bool {
	match, _ := regexp.MatchString("[^@]+@[^@]+\\.[^@]+", mail)
	return match
}
