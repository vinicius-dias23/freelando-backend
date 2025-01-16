package cmd

import "net/http"

func LoadProfile(w http.ResponseWriter, r *http.Request) {
	var p Profile

	err := writeJSON(w, http.StatusOK, p)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func ReturnMyJobs(w http.ResponseWriter, r *http.Request) {
	var j Job

	err := writeJSON(w, http.StatusOK, j)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}
}
