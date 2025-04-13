package main

import (
	// "GolangAdvanced/configs"
	"GolangAdvanced/configs"
	"GolangAdvanced/internal/auth"
	"GolangAdvanced/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	// /auth/login
	// /auth/registration

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening in port 8081")
	server.ListenAndServe()
}
