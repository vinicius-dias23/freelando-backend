package cmd

import "net/http"

func ReturnJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []Job

	err := writeJSON(w, http.StatusOK, jobs)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func SearchJobs(w http.ResponseWriter, r *http.Request) {
	var job Job

	err := readJSON(r, job)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = writeJSON(w, http.StatusOK, job)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func ReturnJob(w http.ResponseWriter, r *http.Request) {
	var job Job

	err := readJSON(r, job)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = writeJSON(w, http.StatusOK, job)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}
}
