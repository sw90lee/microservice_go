package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// Server Define
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Server Start

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
