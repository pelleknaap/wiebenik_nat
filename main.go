package main

import (
	"fmt"
	"github.com/pelleknaap/wiebeniknat_backend/server"
	"net/http"
	"time"
)

func main() {
	sv, err := server.NewServer()
	if err != nil {
		fmt.Printf("An error occurred while initializing the server: %s\n", err)
		return
	}

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: sv.Handler,
		Addr: sv.Config.Address,
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Printf("An error occurred while starting the HTTP server: %s\n", err)
	}
}