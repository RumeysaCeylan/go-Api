package valid

import (
	"fmt"
	"net/http"
)

//check information is valid in sign up
var firstname, lastname, email, password string

func IsValid(w http.ResponseWriter, r *http.Request) bool {
	fNameCheck := IsEmpty(firstname)
	lNameCheck := IsEmpty(lastname)
	emailCheck := IsEmpty(email)
	passwordCheck := IsEmpty(password)
	if fNameCheck || lNameCheck || emailCheck || passwordCheck {
		return true
	} else {
		fmt.Fprintf(w, "EMPTY \n")
		return false
	}
}
func IsEmpty(veri string) bool {
	if len(veri) == 0 {
		return true
	} else {
		return false
	}
}
