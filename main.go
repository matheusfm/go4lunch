package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	app, err := SetupApplication()
	if err != nil {
		log.Fatal(err)
	}
	routers := app.Routers()
	srv := &http.Server{
		Handler:      routers,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
