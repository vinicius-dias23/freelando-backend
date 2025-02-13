package main

import (
	"authentication-service/data"
	"net/http"
)

func (app *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.User

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = validateStruct(requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = app.Repo.Insert(requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusCreated, "Usuário criado com sucesso!")
}
