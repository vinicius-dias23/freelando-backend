package cmd

import (
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Cpf      uint64 `json:"cpf"`
	Number   uint64 `json:"number"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var usr *User
	err := readJSON(r, &usr)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = validateStruct(*usr)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, "Usuário criado com sucesso!")
}
