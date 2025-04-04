package main
import (
	"GolangAdvanced/internal/hello"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening in port 8081")
	server.ListenAndServe()
}
