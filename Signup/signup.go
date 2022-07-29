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
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
	login.Id = 3
	login.FirstName = "Ayşe"
	login.LastName = "KOÇ"
	login.Email = "ak@email.com"
	login.Password = "123456"
	for rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	}
	if login.Email == "" || login.Password == "" {
		fmt.Println("cannot be empty")
	} else {
		if user.Email == login.Email {
			fmt.Fprintf(w, "Email is used")
		} else {
			if CheckEmail(login.Email) && valid.IsValid(w, r) {
				db.Exec("INSERT INTO Userr(firstName,lastName,email,password) VALUES($1,$2,$3,$4)", login.FirstName, login.LastName, login.Email, login.Password)
				peopleByte, _ := json.MarshalIndent(user, "", "\t")

				w.Header().Set("Content-Type", "application/json")

				w.Write(peopleByte)

				defer db.Close()
				_ = json.NewDecoder(r.Body).Decode(&login)
				//login.Id = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
				lg = append(lg, login)
				json.NewEncoder(w).Encode(login)

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
