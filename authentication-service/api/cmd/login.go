package cmd

import "net/http"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Logar(w http.ResponseWriter, r *http.Request) {
	var usr UserLogin
	err := readJSON(r, &usr)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}
}
