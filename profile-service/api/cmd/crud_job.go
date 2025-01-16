package cmd

import "net/http"

func CreateJob(w http.ResponseWriter, r *http.Request) {
	var j Job

	err := readJSON(r, j)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = validateStruct(j)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, "Freela Criado")
}

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	var j Job

	err := readJSON(r, j)
	if err != nil {
		errorJSON(w, err, http.StatusOK)
		return
	}
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	var j Job

	err := readJSON(r, j)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, "Freela deletado do perfil")
}

func InactivateJob(w http.ResponseWriter, r *http.Request) {
	var j Job

	err := readJSON(r, j)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, "Freela desativado do perfil")
}
