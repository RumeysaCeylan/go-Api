package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golib/postgresql"
	"log"
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
	Status  int    `json:"status"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}

var user User
var jsonresp JsonResponse

func Login(w http.ResponseWriter, r *http.Request) {
	openConnention := postgresql.OpenConnention()
	r.ParseForm()
	var users []User
	db := openConnention

	r.ParseForm()
	//_firstName := r.FormValue("firstName")
	//_lastName := r.FormValue("lastName")
	_password := r.FormValue("password")
	_email := r.FormValue("email")
	rows, err := db.Query("SELECT * FROM Userr")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no record founds!!")
			return
		}
		fmt.Println("select error")
		log.Fatal(err)

	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r) // Gets params
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
		}
		users = append(users, user)

	}
	jsonresp.Message = "Login failed"
	jsonresp.Status = 0
	for i := 0; i < len(users); i++ {
		if _email == users[i].Email && _password == users[i].Password {

			jsonresp.Data = append(jsonresp.Data, users[i])
			jsonresp.Message = "Login successful"
			jsonresp.Status = 1

			//json.NewEncoder(w).Encode(users)
			break
		}
	}
	json.NewEncoder(w).Encode(jsonresp.Data)

}

/*for _, item := range users {
	if _email == item.Email || _password == item.Password {

		//if string(item.Id) == params["id"] {
		fmt.Fprintf(w, "Login successful\n")
		fmt.Println(_firstName + " " + _lastName)
		//peopleByte, _ := json.MarshalIndent(user, "", "\t")
		//w.Write(peopleByte)
		jsonresp.Data = append(jsonresp.Data, item)
		jsonresp.Message = "Login successful"
		json.NewEncoder(w).Encode(jsonresp)
		//json.NewEncoder(w).Encode(users)
		return
		//}

	} else {
		fmt.Fprintf(w, "Failed login\n")
	}
}*/ //json.NewEncoder(w).Encode(&User{})
