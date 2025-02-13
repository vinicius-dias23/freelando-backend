package main

import (
	"fmt"
	"log"
	"net/http"
	"profile-service/api"
)

const webPort = "80"

func main() {
	log.Printf("Listening on port :%s", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api.Routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
