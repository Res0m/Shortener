package main

import (
	"GolangAdvanced/configs"
	"GolangAdvanced/internal/hello"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening in port 8081")
	server.ListenAndServe()
}
