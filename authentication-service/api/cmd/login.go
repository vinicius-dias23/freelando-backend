package cmd

import "net/http"

func Logar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login realizado com sucesso!"))
}
