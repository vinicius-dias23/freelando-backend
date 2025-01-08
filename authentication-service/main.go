package main

import (
	"authentication-service/api"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	fmt.Println("Escutando na porta :80")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api.Routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
